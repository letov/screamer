package grab

import (
	"fmt"
	mux "github.com/gorilla/mux"
	"net/http"
	config "screamer/internal/config"
	handlers "screamer/internal/grab/handlers"
)

var hs = []GrabHandler{
	{Route: handlers.UpdateRoute, Handler: handlers.UpdateHandler},
	{Route: handlers.DebugRoute, Handler: handlers.DebugHandler},
}

func Init() {
	c := config.GetConfig()
	router := getRouter()

	addr := fmt.Sprintf(":%v", c.Port)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		panic(err)
	}
}

type GrabHandler struct {
	Route   string
	Handler func(res http.ResponseWriter, req *http.Request)
}

func getRouter() *mux.Router {
	router := mux.NewRouter()
	for _, h := range hs {
		router.HandleFunc(h.Route, h.Handler)
	}
	return router
}
