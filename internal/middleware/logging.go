package middleware

import (
	"log/slog"
	"net/http"
	"os"
	"time"
)

var logger *slog.Logger

func init() {
	options := &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == "url" {
				return slog.String("url", a.String())
			}
			return a
		},
	}

	logger = slog.New(slog.NewJSONHandler(os.Stdout, options))
}

// LoggingMiddleware logs the incoming HTTP request & its duration using structured logging.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		logger.Info("Started request", slog.String("method", r.Method), slog.String("url", r.URL.Path))

		next.ServeHTTP(w, r)

		logger.Info("Completed request", slog.String("url", r.URL.Path), slog.Duration("duration", time.Since(start)))
	})
}
