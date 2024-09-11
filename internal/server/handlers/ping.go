package handlers

import (
	"context"
	"net/http"
	"screamer/internal/server/db"
)

type PingHandler struct {
	db *db.DB
}

func (h *PingHandler) Handler(res http.ResponseWriter, _ *http.Request) {
	if err := h.db.Ping(context.Background()); err == nil {
		res.WriteHeader(http.StatusOK)
		return
	} else {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
}

func NewPingHandler(db *db.DB) *PingHandler {
	return &PingHandler{
		db: db,
	}
}
