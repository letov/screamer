package handlers

import (
	"net/http"
	"screamer/internal/storage"
)

func Debug(res http.ResponseWriter, _ *http.Request) {
	debug := []byte(storage.GetStorage().Debug())
	_, _ = res.Write(debug)
}
