package grab

import (
	"fmt"
	mux "github.com/gorilla/mux"
	"net/http"
	config "screamer/internal/config"
	metric "screamer/internal/metric"
)

type Grab struct{}

func Init() {
	conf := config.NewConfig()
	router := getRouter()

	addr := fmt.Sprintf(":%v", conf.Port)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}

func getRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(`/update/{type:[a-z]+}/{name:[a-z]+}/{value:[-+]?[0-9]*\.?[0-9]+}`, updateHandler)
	return router
}

func updateHandler(res http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	m, err := metric.GetMetric(params["type"], params["name"], params["value"])
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
	}
	fmt.Println(m)
}
