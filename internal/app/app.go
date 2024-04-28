package app

import (
	"context"
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/vpbuyanov/short-url/configs"
	"github.com/vpbuyanov/short-url/internal/handlers"
)

type App interface {
	Start(ctx context.Context) error
}

type app struct {
	cfg    *configs.Config
	logger *logrus.Logger
}

func New(config *configs.Config, log *logrus.Logger) App {
	return &app{
		cfg:    config,
		logger: log,
	}
}

func (app *app) Start(ctx context.Context) error {
	h := handlers.New(app.logger)

	http.HandleFunc("/", h.Shorter)

	return http.ListenAndServe("0.0.0.0:8080", nil)
}
