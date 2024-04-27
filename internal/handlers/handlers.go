package handlers

import (
	"net/http"
)

type Handler interface {
	Shorter(w http.ResponseWriter, r *http.Request)
}

type handlers struct{}

func New() Handler {
	return new(handlers)
}
