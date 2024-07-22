package middleware

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-redis/redis/v8"
	"log/slog"
	"net/http"
	"time"
)

var (
	ctx         = context.Background()
	redisClient *redis.Client
)

func init() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

// GenerateCacheKey is a placeholder function that should be implemented to generate a cache key based on the request URL and query parameters.
func GenerateCacheKey(r *http.Request) string {
	// Return a unique cache key based on the request URL and query parameters
	return r.URL.String()
}

// FetchData is a placeholder function that should be implemented to fetch data from the database or another service.
func FetchData() interface{} {
	return nil
}

func Cache(next http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			err error
		)

		slog.Info("Cache middleware",
			slog.String("method", r.Method),
			slog.String("url", r.URL.String()),
		)

		cacheKey := GenerateCacheKey(r) // Implement this based on your URL and query parameters
		cachedData, err := redisClient.Get(ctx, cacheKey).Result()

		if errors.Is(err, redis.Nil) {
			// Cache miss, fetch data and set it in cache
			data := FetchData() // Implement data fetching logic
			jsonData, _ := json.Marshal(data)
			redisClient.Set(ctx, cacheKey, jsonData, 30*time.Minute) // Adjust expiration as needed

			if _, err = w.Write(jsonData); err != nil {
				return
			}
		} else if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		} else {
			if _, err = w.Write([]byte(cachedData)); err != nil {
				return
			}
		}

		// Add caching logic here
		next.ServeHTTP(w, r)
	}
}
