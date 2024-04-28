package main

import (
	"log/slog"
	"net/http"
)

func logging(logger *slog.Logger, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logger.Info("handling request", "method", r.Method, "path", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
