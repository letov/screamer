package pusher

import (
	"fmt"
	"log"
	"net/http"
	"screamer/internal/collector"
	"screamer/internal/config"
)

func SendData() {
	gauge, counter := collector.GetMetrics()
	for k, v := range *gauge {
		go request("update", "gauge", k, fmt.Sprintf("%f", v))
	}
	for k, v := range *counter {
		go request("update", "counter", k, fmt.Sprintf("%v", v))
	}
}

func request(method, kind, name, value string) {
	c := config.GetConfig()
	url := fmt.Sprintf("%v/%v/%v/%v/%v", c.ServerUrl, method, kind, name, value)
	r, err := http.Post(url, "text/plain", nil)
	if c.AgentLogEnable {
		if err != nil {
			log.Println("Request error ", err.Error())
		} else if r.StatusCode != http.StatusOK {
			log.Println("Bad status ", r.StatusCode)
		}
	}
}
