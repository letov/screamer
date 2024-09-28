package middlewares

import (
	"bytes"
	"io"
	"net/http"
	"screamer/internal/common"
	"screamer/internal/common/hash"
	"screamer/internal/server/config"
)

func CheckHash(c *config.Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if c.Key != "" {
				h := r.Header.Get("HashSHA256")
				reader := io.NopCloser(r.Body)
				body, _ := io.ReadAll(reader)
				_ = r.Body.Close()
				if h != hash.Encode(&body, c.Key) {
					http.Error(w, common.ErrInvalidHash.Error(), http.StatusBadRequest)
					return
				}
				r.Body = io.NopCloser(bytes.NewReader(body))
			}

			next.ServeHTTP(w, r)
		})
	}
}
