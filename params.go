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
