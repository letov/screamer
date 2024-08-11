package config

import (
	"log"
	"strconv"
	"strings"
)

func getPortFromUrl(serverURL string) string {
	d := strings.Split(serverURL, ":")
	if len(d) == 2 {
		_, err := strconv.Atoi(d[1])
		if err == nil {
			return d[1]
		}
	}
	log.Fatal("Fail parse port value")
	return ""
}
