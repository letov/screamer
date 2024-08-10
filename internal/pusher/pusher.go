package pusher

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"screamer/internal/collector"
	"screamer/internal/config"
)

func SendData() {
	for l, e := range collector.Export() {
		for k, v := range e {
			request("update", string(l), k, v)
		}
	}
}

func request(method, kind, name, value string) {
	c := config.GetConfig()
	url := fmt.Sprintf("%v/%v/%v/%v/%v", c.ServerUrl, method, kind, name, value)
	r, err := http.Post(url, "text/plain", nil)
	if err == nil {
		defer func(Body io.ReadCloser) {
			_, _ = io.Copy(io.Discard, r.Body)
			_ = Body.Close()
		}(r.Body)
	}
	if c.AgentLogEnable {
		if err != nil {
			log.Println("Request error", err.Error())
		} else if r.StatusCode != http.StatusOK {
			log.Println("Bad status", r.StatusCode)
		} else {
			data, _ := io.ReadAll(r.Body)
			log.Println("Answer", string(data))
		}
	}
}
