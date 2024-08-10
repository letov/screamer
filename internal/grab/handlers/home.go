package handlers

import (
	"fmt"
	"net/http"
	"screamer/internal/common"
	"screamer/internal/metric"
	"screamer/internal/storage"
)

type LastValues struct {
	Title       string
	MetricNames *[]string
	MetricKind  metric.Kind
}

var lvs = []LastValues{
	{
		Title:       "Counters",
		MetricNames: common.GetCounterInit(),
		MetricKind:  metric.Counter,
	},
	{
		Title:       "Gauges",
		MetricNames: common.GetGaugeInit(),
		MetricKind:  metric.Gauge,
	},
}

func Home(res http.ResponseWriter, _ *http.Request) {
	s := storage.GetStorage()
	if s == nil {
		http.Error(res, ErrNoStorage.Error(), http.StatusBadRequest)
		return
	}

	r := "<html><body>"
	r += fmt.Sprintf("<h1>Metrics</h1>")
	for _, lv := range lvs {
		r += fmt.Sprintf("<h2>%v</h2>", lv.Title)
		for _, mn := range *lv.MetricNames {
			v, err := s.GetLastAsString(lv.MetricKind, mn)
			if err == nil {
				r += fmt.Sprintf("<p>%v: %v</p>", mn, v)
			} else {
				r += fmt.Sprintf("<p>%v: <strong>%v</strong</p>", mn, err.Error())
			}
		}
	}
	r += "</body></html>"

	_, err := res.Write([]byte(r))
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
}
