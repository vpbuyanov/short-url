package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"github.com/vpbuyanov/short-url/internal/helper"
)

type Handler interface {
	RegisterRouter(ctx fiber.Router)
}

type handlers struct {
	logger *logrus.Logger
	url    helper.URL
}

func New(log *logrus.Logger) Handler {
	url := helper.New()

	return &handlers{
		logger: log,
		url:    url,
	}
}

func (h *handlers) RegisterRouter(app fiber.Router) {
	app.Post("/", h.createShortURL)
	app.Get("/:id", h.getFullURL)
}
