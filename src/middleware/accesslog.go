package middleware

import (
	"momonga_blog/internal/logging"
	"net/http"
	"time"
)

func AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now()

		next.ServeHTTP(w, r)

		logging.AccessLogger.Info("Access Log",
			"method", r.Method,
			"path", r.URL.Path,
			"remote", r.RemoteAddr,
			"user_agent", r.UserAgent(),
			"duration", time.Since(startTime).String(),
		)
	})
}