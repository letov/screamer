package middlewares

import (
	"net"
	"net/http"
	"screamer/internal/server/config"

	"go.uber.org/zap"
)

func TrustedSubnet(c *config.Config, log *zap.SugaredLogger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if len(c.TrustedSubnet) != 0 {
				agentIpStr := r.Header.Get("X-Real-IP")
				if len(agentIpStr) == 0 {
					log.Fatal("header X-Real-IP is empty")
				}
				agentIp := net.ParseIP(agentIpStr)
				_, tsNet, err := net.ParseCIDR(c.TrustedSubnet)
				if err != nil {
					log.Fatal("invalid trusted subnet", zap.String("subnet", c.TrustedSubnet))
				}
				if !tsNet.Contains(agentIp) {
					log.Fatal("untrusted subnet", zap.String("subnet", c.TrustedSubnet))
				}
				next.ServeHTTP(w, r)
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}
}
