package handlers

import (
	"fmt"
	"log"
	"net/http"
	"screamer/internal/config"
	"screamer/internal/metric"
	"screamer/internal/storage"
	"sort"
)

type LastValues struct {
	Title      string
	MetricKind metric.Kind
}

var lvs = []LastValues{
	{
		Title:      "Counters",
		MetricKind: metric.Counter,
	},
	{
		Title:      "Gauges",
		MetricKind: metric.Gauge,
	},
}

func Home(res http.ResponseWriter, _ *http.Request) {
	s := storage.GetStorage()
	c := config.GetConfig()
	if s == nil {
		http.Error(res, ErrNoStorage.Error(), http.StatusBadRequest)
		return
	}

	r := "<html><body>"
	r += fmt.Sprintf("<h1>Metrics</h1>")
	for _, lv := range lvs {
		r += fmt.Sprintf("<h2>%v</h2>", lv.Title)
		m, err := s.GetAllLastAsString(lv.MetricKind)
		if err == nil {
			var keys []string
			for mn := range *m {
				keys = append(keys, mn)
			}
			sort.Strings(keys)
			for _, k := range keys {
				r += fmt.Sprintf("<p>%v: %v</p>", k, (*m)[k])
			}
		} else if c.ServerLogEnable {
			log.Println("Get all metrics error:", err.Error())
		}
	}
	r += "</body></html>"

	_, err := res.Write([]byte(r))
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
}
