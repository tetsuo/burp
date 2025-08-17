package main

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"strconv"
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
		http.Error(w, "id cannot be blank; only letters and numbers allowed", http.StatusBadRequest)
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
		http.Error(w, "id cannot be blank; only letters and numbers allowed", http.StatusBadRequest)
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
		http.Error(w, "id cannot be blank; only letters and numbers allowed", http.StatusBadRequest)
		return
	}

	// Model choice (string, <=140 chars); validate against registry
	m := r.FormValue("model")
	if m == "" {
		http.Error(w, "model cannot be blank", http.StatusBadRequest)
		return
	}
	if len(m) > 140 {
		http.Error(w, "model must be <= 140 characters", http.StatusBadRequest)
		return
	}

	model := ChatModel(m)
	if provider, exists := providerFor[model]; !exists {
		http.Error(w, "unknown model", http.StatusBadRequest)
		return
	} else if provider == ChatProviderAnthropic {
		if s.wkr.ac == nil {
			http.Error(w, "model not supported", http.StatusServiceUnavailable)
			return
		}
		params, reason := parseAnthropicParams(r, model)
		if reason != "" {
			http.Error(w, reason, http.StatusBadRequest)
			return
		}
		s.publishMessage(w, r, model, id, params)
	} else if provider == ChatProviderOpenAI {
		if s.wkr.oc == nil {
			http.Error(w, "model not supported", http.StatusServiceUnavailable)
			return
		}
		params, reason := parseOpenAIParams(r, model)
		if reason != "" {
			http.Error(w, reason, http.StatusBadRequest)
			return
		}
		s.publishMessage(w, r, model, id, params)
	}
}

func (s *Server) publishMessage(w http.ResponseWriter, r *http.Request, model ChatModel, id string, params messageParams) {
	b, err := io.ReadAll(http.MaxBytesReader(w, r.Body, modelMaxInputChars[model]))
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
	go s.wkr.Send(context.Background(), id, body, model, params)

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
		http.Error(w, "id cannot be blank; only letters and numbers allowed", http.StatusBadRequest)
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

	var params messageParams

	switch provider {
	case 0:
		if s.wkr.oc != nil {
			model = FallbackOpenAIChatModel // default when both clients exist
			params, _ = parseOpenAIParams(r, model)
		} else {
			model = FallbackAnthropicChatModel
			params, _ = parseAnthropicParams(r, model)
		}
	case ChatProviderAnthropic:
		if s.wkr.ac == nil {
			http.Error(w, "model not supported", http.StatusServiceUnavailable)
			return
		} else if model == "" {
			model = FallbackAnthropicChatModel
		}
		params, _ = parseAnthropicParams(r, model)
	case ChatProviderOpenAI:
		if s.wkr.oc == nil {
			http.Error(w, "model not supported", http.StatusServiceUnavailable)
			return
		} else if model == "" {
			model = FallbackOpenAIChatModel
		}
		params, _ = parseOpenAIParams(r, model)
	}

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
					<span id="params" class="params"></span>
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
	io.WriteString(w, string(model))
	io.WriteString(w, `',
        temperature: `)
	io.WriteString(w, strconv.FormatFloat(params.Temperature, 'g', 17, 64))
	io.WriteString(w, `,
        maxTokens: `)
	io.WriteString(w, strconv.FormatInt(params.MaxTokens, 10))

	if params.TopP != nil {
		io.WriteString(w, `,
        topP: `)
		io.WriteString(w, strconv.FormatFloat(*params.TopP, 'g', 17, 64))
	}

	if params.TopK != nil {
		io.WriteString(w, `,
        topK: `)
		io.WriteString(w, strconv.FormatInt(*params.TopK, 10))
	}

	io.WriteString(w, `,
        subscribeUrl: new URL('/', window.location.href),
        publishUrl: new URL('/', window.location.href),
      });

      chat.mountTo(document);
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
