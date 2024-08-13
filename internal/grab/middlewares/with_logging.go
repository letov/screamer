package middlewares

import (
	"github.com/go-chi/chi/v5/middleware"
	"net/http"
	"screamer/internal/logger"
	"time"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		uri := r.RequestURI
		method := r.Method

		ww := middleware.NewWrapResponseWriter(w, r.ProtoMajor)
		next.ServeHTTP(ww, r)

		duration := time.Since(start)
		s := *logger.GetSugar()

		s.Infoln(
			"uri", uri,
			"method", method,
			"duration", duration,
			"status", ww.Status(),
			"size", ww.BytesWritten(),
		)
	})
}
