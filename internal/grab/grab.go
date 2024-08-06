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
	c := config.GetConfig()
	router := getRouter()

	addr := fmt.Sprintf(":%v", c.Port)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}

func getRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc(`/update/{label:[a-zA-Z]+}/{name:[a-zA-Z]+}/{value:[-+]?[0-9]*\.?[0-9]+}`, updateHandler)
	return router
}

func updateHandler(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(res, "", http.StatusNotFound)
		return
	}
	params := mux.Vars(req)
	m, err := metric.NewMetric(metric.Raw{
		Label: params["label"],
		Name:  params["name"],
		Value: params["value"],
	})
	if err != nil {
		http.Error(res, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Println(m)
}
