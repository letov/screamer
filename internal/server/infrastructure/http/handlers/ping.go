package handlers

import (
	"context"
	"net/http"
	"screamer/internal/server/infrastructure/db"
	"time"
)

type PingHandler struct {
	db *db.DB
}

func (h *PingHandler) Handler(res http.ResponseWriter, _ *http.Request) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := h.db.Ping(ctx); err == nil {
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
