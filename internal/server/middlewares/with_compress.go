package middlewares

import (
	"compress/gzip"
	"io"
	"net/http"
	"slices"
	"strings"
)

type gzipWriter struct {
	http.ResponseWriter
	Types  []string
	Writer io.Writer
}

func (w gzipWriter) Write(b []byte) (int, error) {
	if slices.Contains(w.Types, strings.Split(w.Header().Get("Content-Type"), ":")[0]) {
		w.Header().Set("Content-Encoding", "gzip")
		return w.Writer.Write(b)
	}

	return w.ResponseWriter.Write(b)
}

func Compress(types []string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
				next.ServeHTTP(w, r)
				return
			}

			gz, err := gzip.NewWriterLevel(w, gzip.BestSpeed)
			if err != nil {
				return
			}
			defer func(gz *gzip.Writer) {
				_ = gz.Close()
			}(gz)

			next.ServeHTTP(gzipWriter{Types: types, ResponseWriter: w, Writer: gz}, r)
		})
	}
}