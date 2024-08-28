package pusher

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"screamer/internal/agent/config"
	"screamer/internal/collector"
	"screamer/internal/collector/maps"
)

func SendData() {
	for _, jsms := range collector.ExportJsonMetrics() {
		for _, jsm := range jsms {
			request("update", jsm)
		}
	}
}

func request(method string, jsm *maps.JsonMetric) {
	c := config.GetConfig()
	url := fmt.Sprintf("%v/%v", c.NetAddress.String(), method)
	body, _ := json.Marshal(&jsm)

	r, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err == nil {
		defer func(Body io.ReadCloser) {
			_, _ = io.Copy(io.Discard, r.Body)
			_ = Body.Close()
		}(r.Body)
	}

	if *c.AgentLogEnable {
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
