package main

import "net/http"

func corsHandler(next http.Handler) http.Handler {
	const (
		allowOrigin  = "*"
		allowMethods = "GET, HEAD, POST, OPTIONS"
		allowHeaders = "Content-Type"
	)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		w.Header().Set("Access-Control-Allow-Methods", allowMethods)
		w.Header().Set("Access-Control-Allow-Headers", allowHeaders)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}
		next.ServeHTTP(w, r)
	})
}
