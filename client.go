package main

import (
	"context"
	_ "embed"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/openai/openai-go/v2"
	"github.com/tetsuo/bbq"
)

type Worker struct {
	c openai.Client
}

func NewWorker(oc openai.Client) *Worker {
	return &Worker{c: oc}
}

//go:embed prompt.md
var systemMsg string

func (w *Worker) Send(ctx context.Context, id, userMsg string) {
	q := bbq.New[string](16)

	go func(q *bbq.BBQ[string]) {
		// pull last entries
		hist := snapshotHistory(id, keepMin)

		msgs := []openai.ChatCompletionMessageParamUnion{
			openai.SystemMessage(systemMsg),
		}
		msgs = append(msgs, historyToMessages(hist)...)

		params := openai.ChatCompletionNewParams{
			Model:    openai.ChatModelGPT5Nano,
			Messages: msgs,
		}

		stream := w.c.Chat.Completions.NewStreaming(ctx, params)
		defer stream.Close()

		for stream.Next() {
			q.Write(stream.Current().Choices[0].Delta.Content)
		}
		if err := stream.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "\nError: %v\n", err)
		}
		q.Close()
	}(q)

	go func(q *bbq.BBQ[string]) {
		// emit a batch when either 10 tokens is reached, or 5 seconds have passed
		for batch := range q.SlicesWhen(10, time.Second*5) {
			if len(batch) < 1 {
				continue
			}
			body := strings.Join(batch, "")
			if len(strings.TrimSpace(body)) == 0 {
				continue
			}
			publish(&Message{
				ID:   id,
				Role: AssistantMessage,
				Body: body,
			})
		}
		publish(&Message{
			ID:   id,
			Role: AssistantMessage,
			Body: "",
		})
	}(q)
}

// snapshotHistory copies recent messages for a channel.
func snapshotHistory(id string, max int) []*Message {
	mu.Lock()
	list := recent[id]
	// copy to avoid races
	cp := make([]*Message, len(list))
	for i := range list {
		cp[i] = list[i].Message
	}
	mu.Unlock()

	// keep only last max
	start := 0
	if max > 0 && len(cp) > max {
		start = len(cp) - max
	}
	out := make([]*Message, 0, len(cp)-start)
	for _, msg := range cp[start:] {
		if msg.LongPollTimeout ||
			msg.Body == "" ||
			msg.Role != UserMessage && msg.Role != AssistantMessage {
			continue
		}
		out = append(out, msg)
	}
	return out
}

func historyToMessages(msgs []*Message) []openai.ChatCompletionMessageParamUnion {
	p := make([]openai.ChatCompletionMessageParamUnion, 0, len(msgs))
	for _, msg := range msgs {
		switch msg.Role {
		case UserMessage:
			p = append(p, openai.UserMessage(msg.Body))
		case AssistantMessage:
			p = append(p, openai.AssistantMessage(msg.Body))
		}
	}
	return p
}
