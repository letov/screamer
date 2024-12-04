package middlewares

import (
	"net/http"
	"screamer/internal/common/logger"
	"time"

	"github.com/go-chi/chi/v5/middleware"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		uri := r.RequestURI
		method := r.Method

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)

		duration := time.Since(start)
		s := logger.NewLogger()

		s.Infoln(
			"uri", uri,
			"method", method,
			"duration", duration,
			"status", ww.Status(),
			"size", ww.BytesWritten(),
		)
	})
}
