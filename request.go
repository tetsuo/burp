package main

import (
	"net/http"
	"strconv"
)

type messageParams struct {
	// always set
	MaxTokens   int64
	Temperature float64
	// optionals
	TopP *float64
	TopK *int64 // Anthropic only; always nil for OpenAI
}

func (s *Server) parseRequest(r *http.Request) (id string, model ChatModel, provider ChatProvider, params messageParams, reason string) {
	id, reason = parseID(r)
	if reason != "" {
		return
	}

	model, provider, reason = parseModel(r)
	if reason != "" {
		return
	}

	switch provider {
	case 0:
		reason = "model not supported"
	case ChatProviderAnthropic:
		if s.wkr.ac == nil {
			reason = "model not supported"
			return
		}
		params, reason = parseAnthropicParams(r, model)
		if reason != "" {
			return
		}
	case ChatProviderOpenAI:
		if s.wkr.oc == nil {
			reason = "model not supported"
			return
		}
		params, reason = parseOpenAIParams(r, model)
		if reason != "" {
			return
		}
	}
	return
}

func parseID(r *http.Request) (string, string) {
	id := r.FormValue("id")
	if id == "" {
		return "", "id cannot be blank"
	}
	if len(id) > 32 {
		return "", "id must be <= 32 characters"
	}
	if !isNonEmptyAlnum(id) {
		return "", "id must be alphanumeric"
	}
	return id, ""
}

func parseModel(r *http.Request) (ChatModel, ChatProvider, string) {
	m := r.FormValue("model")
	if m == "" {
		return "", 0, "model cannot be blank"
	}
	if len(m) > 140 {
		return "", 0, "model must be <= 140 characters"
	}
	model := ChatModel(m)
	return model, providerFor[model], ""
}

func parseOpenAIParams(r *http.Request, model ChatModel) (params messageParams, reason string) {
	limit, ok := modelMaxOutputTokens[model]
	if !ok || limit <= 0 {
		panic("unknown model or token limit not configured")
	}

	// temperature: [0.0, 2.0], default 1.0
	params.Temperature = 1.0
	if s := r.FormValue("temp"); s != "" {
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			reason = "temp must be a number"
		} else if v < 0.0 || v > 2.0 {
			reason = "temp out of range for OpenAI (0.0–2.0)"
		} else {
			params.Temperature = v
		}
	}

	// top_p: [0.0, 1.0], default 1.0
	if s := r.FormValue("top_p"); s != "" {
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			reason = "top_p must be a number"
		} else if v < 0.0 || v > 1.0 {
			reason = "top_p out of range (0.0–1.0)"
		} else {
			params.TopP = &v
		}
	}

	params.MaxTokens = limit
	if s := r.FormValue("max_tokens"); s != "" {
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			reason = "max_tokens must be an integer"
		} else if v < 0 || v > limit {
			reason = "max_tokens out of range"
		} else {
			params.MaxTokens = v
		}
	}

	return
}

func parseAnthropicParams(r *http.Request, model ChatModel) (params messageParams, reason string) {
	limit, ok := modelMaxOutputTokens[model]
	if !ok || limit <= 0 {
		panic("unknown model or token limit not configured")
	}

	// temperature: [0.0, 1.0], default 1.0
	params.Temperature = 1.0
	if s := r.FormValue("temp"); s != "" {
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			reason = "temp must be a number"
		} else if v < 0.0 || v > 1.0 {
			reason = "temp out of range for Anthropic (0.0–1.0)"
		} else {
			params.Temperature = v
		}
	}

	// top_p: [0.0, 1.0], default nil (0.99)
	if s := r.FormValue("top_p"); s != "" {
		v, err := strconv.ParseFloat(s, 64)
		if err != nil {
			reason = "top_p must be a number"
		} else if v < 0.0 || v > 1.0 {
			reason = "top_p out of range (0.0–1.0)"
		} else {
			params.TopP = &v
		}
	}

	// top_k: [0, 500], default nil (40)
	if s := r.FormValue("top_k"); s != "" {
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			reason = "top_k must be an integer"
		} else if v < 0 || v > 500 {
			reason = "top_k out of range for Anthropic (0–500)"
		} else {
			params.TopK = &v
		}
	}

	params.MaxTokens = limit
	if s := r.FormValue("max_tokens"); s != "" {
		v, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			reason = "max_tokens must be an integer"
		} else if v < 0 || v > limit {
			reason = "max_tokens out of range"
		} else {
			params.MaxTokens = v
		}
	}

	return
}

func isNonEmptyAlnum(s string) bool {
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
