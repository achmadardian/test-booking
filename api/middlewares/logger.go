package middlewares

import (
	"log/slog"
	"net/http"
	"time"
)

var logger *slog.Logger

func SetLogger(l *slog.Logger) {
	logger = l
}

func Logger(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		logger.Info("request",
			slog.String("path", r.URL.Path),
			slog.String("method", r.Method),
		)

		h.ServeHTTP(w, r)

		d := time.Since(start)
		logger.Info("response",
			slog.String("path", r.URL.Path),
			slog.String("method", r.Method),
			slog.Duration("duration", d),
		)
	})
}
