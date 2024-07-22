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
		AddSource: false, // If this is set to true the function and file will be displayed in the log output.
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

type statusRecorder struct {
	http.ResponseWriter
	statusCode int
}

func newStatusRecorder(w http.ResponseWriter) *statusRecorder {
	return &statusRecorder{w, http.StatusOK}
}

func (r *statusRecorder) WriteHeader(code int) {
	r.statusCode = code
	r.ResponseWriter.WriteHeader(code)
}

// LoggingMiddleware logs the incoming HTTP request & its duration using structured logging.
func LoggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		rec := newStatusRecorder(w)
		queryString := r.URL.RawQuery

		if queryString != "" {
			logger.Info("Started request",
				slog.String("method", r.Method),
				slog.String("url", r.URL.Path),
				slog.String("query", queryString)) // Include the query string in the log output

		} else {
			logger.Info("Started request",
				slog.String("method", r.Method),
				slog.String("url", r.URL.Path))
		}

		next.ServeHTTP(rec, r)

		if queryString != "" {
			logger.Info("Completed request",
				slog.String("method", r.Method),
				slog.String("url", r.URL.Path),
				slog.String("query", queryString),
				slog.Int("status", rec.statusCode),
				slog.Duration("duration", time.Duration(time.Since(start).Milliseconds())))
		} else {
			logger.Info("Completed request",
				slog.String("method", r.Method),
				slog.String("url", r.URL.Path),
				slog.Int("status", rec.statusCode),
				slog.Duration("duration", time.Duration(time.Since(start).Milliseconds())))
		}
	})
}
