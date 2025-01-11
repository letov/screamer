package middlewares

import (
	"net"
	"net/http"
	"screamer/internal/common"
	"screamer/internal/server/infrastructure/config"
)

func TrustedSubnet(c *config.Config) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if len(c.TrustedSubnet) != 0 {
				agentIPStr := r.Header.Get("X-Real-IP")
				if len(agentIPStr) == 0 {
					http.Error(w, common.ErrXRealIPEmpty.Error(), http.StatusBadRequest)
					return
				}
				agentIP := net.ParseIP(agentIPStr)
				_, tsNet, err := net.ParseCIDR(c.TrustedSubnet)
				if err != nil {
					http.Error(w, common.ErrTrustedSubnetInvalid.Error(), http.StatusBadRequest)
					return
				}
				if !tsNet.Contains(agentIP) {
					http.Error(w, common.ErrTrustedSubnetUnTrust.Error(), http.StatusBadRequest)
					return
				}
				next.ServeHTTP(w, r)
			} else {
				next.ServeHTTP(w, r)
			}
		})
	}
}
