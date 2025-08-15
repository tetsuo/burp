package main

import (
	"encoding/json"
	"log"
	"sync"
	"time"

	"go4.org/types"
)

type MessageRole uint8

const (
	SystemMessage MessageRole = iota
	AssistantMessage
	UserMessage
)

type Message struct {
	// ID is the channel ID.
	ID string `json:",omitempty"`
	// Body is the input.
	Body string
	// Time is the time the message was received, or the time of the
	// long poll timeout. This is what clients should send as the
	// "after" URL parameter for the next message.
	Time types.Time3339
	// LongPollTimeout indicates that no message received and the
	// client should retry with ?after=<Time>.
	LongPollTimeout bool `json:",omitempty"`
	// Role of the message sent.
	Role MessageRole `json:",omitempty"`
}

type messageAndJSON struct {
	*Message
	json string // JSON MarshalIndent of Message
}

var (
	mu      sync.Mutex                                       // guards following
	recent  = map[string][]*messageAndJSON{}                 // newest at end
	waiting = map[string]map[chan *messageAndJSON]struct{}{} // long-poll chans
)

func newMessageAndJSON(msg *Message) *messageAndJSON {
	msg.Time = types.Time3339(time.Now())
	j, err := json.MarshalIndent(msg, "", "\t")
	if err != nil {
		log.Printf("JSON marshal error: %v", err)
	}
	return &messageAndJSON{Message: msg, json: string(j)}
}

const keepMin = 50

// trimRecentLocked trims the per-ID ring buffer in-place.
//
// Must be called with mu held.
func trimRecentLocked(id string) {
	const maxAge = 1 * time.Hour

	list := recent[id]
	if len(list) <= keepMin {
		return
	}

	cutoff := time.Now().Add(-maxAge)

	trim := 0
	for trim < len(list) &&
		len(list)-trim > keepMin && // never trim below keepMin
		list[trim].Time.Time().Before(cutoff) {
		trim++
	}
	if trim == 0 {
		return // nothing to do
	}

	// shift the tail left, keep the same backing array
	copy(list, list[trim:])
	recent[id] = list[:len(list)-trim]
}

func register(id string, ch chan *messageAndJSON, after time.Time) {
	mu.Lock()
	defer mu.Unlock()

	for _, msg := range recent[id] {
		if msg.Time.Time().After(after) {
			ch <- msg
			return
		}
	}

	if waiting[id] == nil {
		waiting[id] = make(map[chan *messageAndJSON]struct{})
	}
	waiting[id][ch] = struct{}{}
}

func unregister(id string, ch chan *messageAndJSON) {
	mu.Lock()
	defer mu.Unlock()
	delete(waiting[id], ch)
	if len(waiting[id]) == 0 {
		delete(waiting, id) // hygiene
	}
}

func publish(msg *Message) {
	if msg.ID == "" {
		log.Printf("message dropped: missing channel ID")
		return
	}

	mj := newMessageAndJSON(msg)

	mu.Lock()
	defer mu.Unlock()

	recent[msg.ID] = append(recent[msg.ID], mj)
	trimRecentLocked(msg.ID)

	for ch := range waiting[msg.ID] {
		ch <- mj
		delete(waiting[msg.ID], ch)
	}
	if len(waiting[msg.ID]) == 0 {
		delete(waiting, msg.ID)
	}
}
