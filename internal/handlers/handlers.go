package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"github.com/vpbuyanov/short-url/configs"
	"github.com/vpbuyanov/short-url/internal/services"
)

type Handler interface {
	RegisterRouter(ctx fiber.Router)
}

type handlers struct {
	logger *logrus.Logger
	url    services.URL
	cfg    *configs.Server
}

func New(log *logrus.Logger, cfg *configs.Server) Handler {
	url := services.New()

	return &handlers{
		logger: log,
		url:    url,
		cfg:    cfg,
	}
}

func (h *handlers) RegisterRouter(app fiber.Router) {
	app.Post("/", h.createShortURL)
	app.Get("/:id", h.getFullURL)
}
