package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"github.com/vpbuyanov/short-url/internal/configs"
	"github.com/vpbuyanov/short-url/internal/usecase"
)

type Handlers struct {
	logger *logrus.Logger
	url    *usecase.URL
	cfg    *configs.Server
}

func New(cfg *configs.Server, url *usecase.URL, log *logrus.Logger) Handlers {
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
