package custommiddleware

import (
	"context"
	"github.com/richmondgoh8/boilerplate/pkg/logger"
	"net/http"
	"time"
)

type responseWriter struct {
	http.ResponseWriter
	status      int
	wroteHeader bool
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriter {
	return &responseWriter{ResponseWriter: w}
}

func (rw *responseWriter) Status() int {
	return rw.status
}

func (rw *responseWriter) WriteHeader(code int) {
	if rw.wroteHeader {
		return
	}

	rw.status = code
	rw.ResponseWriter.WriteHeader(code)
	rw.wroteHeader = true

	return
}

// LoggingMiddleware logs the incoming HTTP request & its duration.
func LoggingMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		fn := func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				if err := recover(); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					logger.Error(err.(string), context.Background(), nil)
				}
			}()
			start := time.Now()
			wrapped := wrapResponseWriter(w)
			next.ServeHTTP(wrapped, r)

			statisticsMap := map[string]interface{}{
				"status":   wrapped.status,
				"method":   r.Method,
				"path":     r.URL,
				"duration": time.Since(start).Milliseconds(),
			}

			logger.Info("finish baking Chicken", r.Context(), statisticsMap)
			//logger.Log(
			//	"status", wrapped.status,
			//	"method", r.Method,
			//	"path", r.URL.EscapedPath(),
			//	"duration", time.Since(start),
			//)
		}

		return http.HandlerFunc(fn)
	}
}

func LoggingMiddlewareV2(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		//ctx := r.Context()
		//ctx = context.WithValue(ctx, trackingCtxKey.name, uuid.GenerateTrackingID())
		start := time.Now()
		wrapped := wrapResponseWriter(w)
		next.ServeHTTP(w, r)
		statisticsMap := map[string]interface{}{
			"status":   wrapped.status,
			"method":   r.Method,
			"path":     r.URL.EscapedPath(),
			"duration": time.Since(start).Milliseconds(),
		}

		logger.Info("finish baking Chicken", r.Context(), statisticsMap)
	}
	return http.HandlerFunc(fn)
}
