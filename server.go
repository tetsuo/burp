package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/tetsuo/burp/static"
)

type Server struct {
	wkr *Worker
}

func (s *Server) serveWait(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	if !isNonEmptyAlnum(id) {
		http.Error(w, "id cannot be empty; only letters and numbers allowed", http.StatusBadRequest)
		return
	}

	var after time.Time
	if v := r.FormValue("after"); v != "" {
		var err error
		after, err = time.Parse(time.RFC3339Nano, v)
		if err != nil {
			http.Error(w, "after must be RFC3339Nano", http.StatusBadRequest)
			return
		}
	} else {
		after = time.Now()
	}

	ch := make(chan *messageAndJSON, 1)
	register(id, ch, after)
	defer unregister(id, ch)

	ctx := r.Context()
	timer := time.NewTimer(30 * time.Second)
	defer timer.Stop()

	var msg *messageAndJSON
	select {
	case <-ctx.Done():
		return
	case <-timer.C:
		msg = newMessageAndJSON(&Message{LongPollTimeout: true})
	case msg = <-ch:
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	io.WriteString(w, msg.json)
}

func serveRecent(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	id := r.FormValue("id")
	if !isNonEmptyAlnum(id) {
		http.Error(w, "id cannot be empty; only letters and numbers allowed", http.StatusBadRequest)
		return
	}

	var after time.Time
	if v := r.FormValue("after"); v != "" {
		var err error
		after, err = time.Parse(time.RFC3339Nano, v)
		if err != nil {
			http.Error(w, "after must be RFC3339Nano", http.StatusBadRequest)
			return
		}
	}

	mu.Lock()

	var buf bytes.Buffer
	buf.WriteString("[\n")
	n := 0
	for i := len(recent[id]) - 1; i >= 0; i-- {
		msg := recent[id][i]
		if msg.Time.Time().Before(after) {
			continue
		}
		if n > 0 {
			buf.WriteString(",\n")
		}
		buf.WriteString(msg.json)
		n++
	}
	mu.Unlock()

	buf.WriteString("\n]\n")
	w.Write(buf.Bytes())
}

func (s *Server) serveAsk(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	if !isNonEmptyAlnum(id) {
		http.Error(w, "id cannot be empty; only letters and numbers allowed", http.StatusBadRequest)
		return
	}

	// Model choice (string, <=140 chars); validate against registry
	m := r.FormValue("model")
	if m == "" {
		http.Error(w, "model cannot be empty", http.StatusBadRequest)
		return
	}
	if len(m) > 140 {
		http.Error(w, "model must be <= 140 characters", http.StatusBadRequest)
		return
	}

	model := ChatModel(m)
	if provider, exists := providerFor[model]; !exists {
		http.Error(w, "unrecognized model", http.StatusBadRequest)
		return
	} else if provider == ChatProviderAnthropic && s.wkr.ac == nil {
		http.Error(w, "model not supported", http.StatusServiceUnavailable)
		return
	} else if provider == ChatProviderOpenAI && s.wkr.oc == nil {
		http.Error(w, "model not supported", http.StatusServiceUnavailable)
		return
	}

	b, err := io.ReadAll(http.MaxBytesReader(w, r.Body, 32<<10))
	if err != nil {
		w.WriteHeader(http.StatusRequestEntityTooLarge)
		return
	}

	body := string(b)

	publish(&Message{
		ID:   id,
		Body: body,
		Role: UserMessage,
	})

	// kick off work in the background
	go s.wkr.Send(context.Background(), id, body, model)

	w.WriteHeader(http.StatusAccepted)
}

func (s *Server) serveRoot(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	io.WriteString(w, `<html><body><h1>chathelper</h1>
<ul>
	<li><b><a href="/chat">/chat</a></b>: chat in a channel</li>
	<li><b><a href="/wait">/wait</a></b>: long-poll 30s for next message (use ?id=&lt;channel&gt;&amp;after=&lt;RFC3339Nano&gt;)</li>
	<li><b><a href="/recent">/recent</a></b>: recent messages in a channel</li>
</ul></body></html>`)
}

func (s *Server) serveChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	id := r.FormValue("id")
	if !isNonEmptyAlnum(id) {
		http.Error(w, "id cannot be empty; only letters and numbers allowed", http.StatusBadRequest)
		return
	}

	var (
		model    ChatModel
		provider ChatProvider
	)

	// Model choice (string, <=140 chars); validate against registry
	m := r.FormValue("model")
	if m != "" && len(m) <= 140 {
		model = ChatModel(m)
		var exists bool
		if provider, exists = providerFor[model]; !exists {
			model = ""
		}
	}

	switch provider {
	case 0:
		model = FallbackOpenAIChatModel
		if s.wkr.ac == nil {
			model = FallbackAnthropicChatModel
		}
		http.Redirect(w, r, fmt.Sprintf("/chat?id=%s&model=%s", id, string(model)), http.StatusFound)
		return
	case ChatProviderAnthropic:
		if s.wkr.ac == nil {
			http.Error(w, "model not supported", http.StatusServiceUnavailable)
			return
		} else if model == "" {
			model = FallbackAnthropicChatModel
			http.Redirect(w, r, fmt.Sprintf("/chat?id=%s&model=%s", id, string(model)), http.StatusFound)
			return
		}
	case ChatProviderOpenAI:
		if s.wkr.oc == nil {
			http.Error(w, "model not supported", http.StatusServiceUnavailable)
			return
		} else if model == "" {
			model = FallbackOpenAIChatModel
			http.Redirect(w, r, fmt.Sprintf("/chat?id=%s&model=%s", id, string(model)), http.StatusFound)
			return
		}
	}

	m = string(model)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	io.WriteString(w, `<!DOCTYPE html>
<html>
  <head>
    <meta charset="utf-8">
    <title></title>
    <link rel="stylesheet" href="/static/frontend/styles.css">
    <script src="/static/frontend/Chat.js"></script>
  </head>
  <body>
    <div class="chat">
      <div class="lines" id="chat-lines"></div>

      <div class="cli">
        <div class="info">
          <span id="time"></span>
          <span id="user"></span>
          <span id="channel"></span>
          <span id="spinner" class="spinner" aria-live="polite"></span>
        </div>

        <form class="input" id="chat-form">
          <input type="text" name="text" id="chat-input" autocomplete="off" autofocus>
        </form>
      </div>
    </div>

    <script>
      const chat = new Chat({
        channel: '`)
	io.WriteString(w, id)
	io.WriteString(w, `',
        model: '`)
	io.WriteString(w, m)
	io.WriteString(w, `',
        subscribeUrl: new URL('/', window.location.href),
        publishUrl: new URL('/', window.location.href),
      })

      chat.mountTo(document)
    </script>
  </body>
</html>`)
}

func (s *Server) Install(mux *http.ServeMux) {
	mux.Handle("/static/",
		http.StripPrefix("/static",
			http.FileServer(http.FS(static.FS)),
		),
	)

	mux.HandleFunc("/", s.serveRoot)
	mux.HandleFunc("/wait", s.serveWait)
	mux.HandleFunc("/recent", serveRecent)
	mux.HandleFunc("/chat", s.serveChat)
	mux.HandleFunc("/ask", s.serveAsk)
}

func isNonEmptyAlnum(s string) bool {
	if s == "" {
		return false
	}
	for i := 0; i < len(s); i++ {
		c := s[i]
		if !(c >= 'a' && c <= 'z') &&
			!(c >= 'A' && c <= 'Z') &&
			!(c >= '0' && c <= '9') {
			return false
		}
	}
	return true
}
