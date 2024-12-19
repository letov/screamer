package middlewares

import (
	"bytes"
	"io"
	"net/http"
	"screamer/internal/common/hash"
	"screamer/internal/server/config"

	"go.uber.org/zap"
)

func Decrypt(c *config.Config, log *zap.SugaredLogger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if len(c.CryptoKey) != 0 {
				reader := io.NopCloser(r.Body)
				body, _ := io.ReadAll(reader)
				_ = r.Body.Close()
				d := hash.NewRSARSADecrypt(c.CryptoKey, log)
				decoded, err := d.Decrypt(body)
				if err != nil {
					log.Warn("decrypt error", zap.Error(err))
				}
				r.Body = io.NopCloser(bytes.NewReader(decoded))
				next.ServeHTTP(w, r)
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}
}
