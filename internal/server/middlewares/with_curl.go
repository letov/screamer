package middlewares

import (
	"net/http"
	"screamer/internal/common/logger"

	"github.com/aoliveti/curling"
)

func Curl(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cmd, _ := curling.NewFromRequest(r)
		s := logger.NewLogger()
		s.Infoln(cmd)
		next.ServeHTTP(w, r)
	})
}
