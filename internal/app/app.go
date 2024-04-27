package app

import (
	"net/http"

	"github.com/vpbuyanov/short-url/internal/handlers"
)

type App interface {
	Start() error
}

type app struct{}

func New() App {
	return new(app)
}

func (app *app) Start() error {
	h := handlers.New()

	http.HandleFunc("/", h.Shorter)

	return http.ListenAndServe("0.0.0.0:8801", nil)
}
