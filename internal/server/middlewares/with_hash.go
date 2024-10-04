package middlewares

import (
	"bytes"
	"io"
	"net/http"
	"screamer/internal/common"
	"screamer/internal/common/hash"
	"screamer/internal/server/config"
)

type hashWriter struct {
	http.ResponseWriter
	c *config.Config
}

func (w hashWriter) Write(b []byte) (int, error) {
	w.Header().Set("HashSHA256", hash.Encode(&b, w.c.Key))
	return w.ResponseWriter.Write(b)
}

func CheckHash(c *config.Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h := r.Header.Get("HashSHA256")
			if c.Key != "" && h != "" {
				reader := io.NopCloser(r.Body)
				body, _ := io.ReadAll(reader)
				_ = r.Body.Close()
				if h != hash.Encode(&body, c.Key) {
					http.Error(w, common.ErrInvalidHash.Error(), http.StatusBadRequest)
					return
				}
				r.Body = io.NopCloser(bytes.NewReader(body))
				next.ServeHTTP(hashWriter{ResponseWriter: w, c: c}, r)
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}
}
