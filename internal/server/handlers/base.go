package handlers

import "net/http"

type HandlerFunc = func(res http.ResponseWriter, req *http.Request)
