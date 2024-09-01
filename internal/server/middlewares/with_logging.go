package middlewares

import (
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
	"net/http"
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
		s := getSugarLogger()

		s.Infoln(
			"uri", uri,
			"method", method,
			"duration", duration,
			"status", ww.Status(),
			"size", ww.BytesWritten(),
		)
	})
}

func getSugarLogger() *zap.SugaredLogger {
	var sugar zap.SugaredLogger

	logger, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	defer func(logger *zap.Logger) {
		_ = logger.Sync()
	}(logger)
	sugar = *logger.Sugar()

	return &sugar
}
