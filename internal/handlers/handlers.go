package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"github.com/vpbuyanov/short-url/internal/configs"
	"github.com/vpbuyanov/short-url/internal/repos"
)

type Handlers struct {
	logger *logrus.Logger
	url    repos.URL
	cfg    *configs.Server
}

func New(log *logrus.Logger, cfg *configs.Server, url repos.URL) Handlers {
	return Handlers{
		logger: log,
		url:    url,
		cfg:    cfg,
	}
}

func (h *Handlers) RegisterRouter(app fiber.Router) {
	app.Post("/", h.createShortURL)
	app.Get("/:id", h.getFullURL)
}
