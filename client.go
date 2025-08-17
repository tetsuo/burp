package main

import (
	"context"
	_ "embed"
	"log"
	"strings"
	"time"

	anthropic "github.com/anthropics/anthropic-sdk-go"
	"github.com/openai/openai-go/v2"
	"github.com/tetsuo/bbq"
)

type Worker struct {
	oc *openai.Client
	ac *anthropic.Client
}

func NewWorker(oc *openai.Client, ac *anthropic.Client) *Worker {
	return &Worker{oc: oc, ac: ac}
}

//go:embed prompt.md
var systemMsg string

func (w *Worker) Send(ctx context.Context, id, userMsg string, model ChatModel, extras messageParams) {
	q := bbq.New[string](16)

	go func(q *bbq.BBQ[string]) {
		switch providerFor[model] {
		case ChatProviderOpenAI:
			if w.oc == nil {
				log.Println("warn: openai not configured; aborting message")
				q.Close()
				return
			}
			w.streamOpenAI(ctx, id, model, extras, q)
		case ChatProviderAnthropic:
			if w.ac == nil {
				log.Println("warn: anthropic not configured; aborting message")
				q.Close()
				return
			}
			w.streamAnthropic(ctx, id, model, extras, q)
		default:
			log.Println("warn: unrecognized model; aborting message")
			q.Close()
			return
		}
	}(q)

	// Batcher: publish chunks and final empty-string terminator:
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
				ID:    id,
				Role:  AssistantMessage,
				Body:  body,
				Model: model,
			})
		}
		// empty-string terminator
		publish(&Message{
			ID:    id,
			Role:  AssistantMessage,
			Body:  "",
			Model: model,
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

func historyToOpenAI(msgs []*Message) []openai.ChatCompletionMessageParamUnion {
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

func (w *Worker) streamOpenAI(ctx context.Context, id string, model ChatModel, extraParams messageParams, q *bbq.BBQ[string]) {
	// pull last entries
	hist := snapshotHistory(id, keepMin)

	msgs := []openai.ChatCompletionMessageParamUnion{
		openai.SystemMessage(systemMsg),
	}
	msgs = append(msgs, historyToOpenAI(hist)...)

	params := openai.ChatCompletionNewParams{
		Model:               openai.ChatModel(model),
		Messages:            msgs,
		MaxCompletionTokens: openai.Int(extraParams.MaxTokens),
		Temperature:         openai.Float(extraParams.Temperature),
	}

	if extraParams.TopP != nil {
		params.TopP = openai.Float(*extraParams.TopP)
	}

	stream := w.oc.Chat.Completions.NewStreaming(ctx, params)
	defer stream.Close()

	for stream.Next() {
		q.Write(stream.Current().Choices[0].Delta.Content)
	}
	if err := stream.Err(); err != nil {
		log.Printf("error: streamOpenAI: %v\n", err)
	}
	q.Close()
}

func (w *Worker) streamAnthropic(ctx context.Context, id string, model ChatModel, extras messageParams, q *bbq.BBQ[string]) {
	// Convert history to anthropic messages
	hist := snapshotHistory(id, keepMin)
	msgs := make([]anthropic.MessageParam, 0, len(hist)+1)
	for _, m := range hist {
		switch m.Role {
		case UserMessage:
			msgs = append(msgs, anthropic.NewUserMessage(anthropic.NewTextBlock(m.Body)))
		case AssistantMessage:
			msgs = append(msgs, anthropic.NewAssistantMessage(anthropic.NewTextBlock(m.Body)))
		}
	}

	params := anthropic.MessageNewParams{
		Model:       anthropic.Model(model),
		MaxTokens:   extras.MaxTokens,
		Temperature: anthropic.Float(extras.Temperature),
		Messages:    msgs,
		System:      []anthropic.TextBlockParam{{Text: systemMsg}},
	}

	if extras.TopP != nil {
		params.TopP = anthropic.Float(*extras.TopP)
	}

	if extras.TopK != nil {
		params.TopK = anthropic.Int(*extras.TopK)
	}

	stream := w.ac.Messages.NewStreaming(ctx, params)
	defer stream.Close()

	for stream.Next() {
		ev := stream.Current()
		switch any := ev.AsAny().(type) {
		case anthropic.ContentBlockDeltaEvent:
			if td, ok := any.Delta.AsAny().(anthropic.TextDelta); ok {
				q.Write(td.Text)
			}
		}
	}
	if err := stream.Err(); err != nil {
		log.Printf("error: streamAnthropic: %v\n", err)
	}
	q.Close()
}
