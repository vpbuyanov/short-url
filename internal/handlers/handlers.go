package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"github.com/vpbuyanov/short-url/internal/services"
)

type Handler interface {
	RegisterRouter(ctx fiber.Router)
}

type handlers struct {
	logger *logrus.Logger
	url    services.URL
}

func New(log *logrus.Logger) Handler {
	url := services.New()

	return &handlers{
		logger: log,
		url:    url,
	}
}

func (h *handlers) RegisterRouter(app fiber.Router) {
	app.Post("/", h.createShortURL)
	app.Get("/:id", h.getFullURL)
}
