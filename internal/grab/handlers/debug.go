package handlers

import (
	"net/http"
	"screamer/internal/storage"
)

const DebugRoute = `/debug`

func DebugHandler(res http.ResponseWriter, _ *http.Request) {
	debug := []byte(storage.GetStorage().Debug())
	_, _ = res.Write(debug)
}
