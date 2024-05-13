package handlers

import (
	"github.com/gofiber/fiber/v2"

	"github.com/vpbuyanov/short-url/internal/configs"
	"github.com/vpbuyanov/short-url/internal/usecase"
)

type Handlers struct {
	url *usecase.URL
	cfg *configs.Server
}

func New(cfg *configs.Server, url *usecase.URL) Handlers {
	return Handlers{
		url: url,
		cfg: cfg,
	}
}

func (h *Handlers) RegisterRouter(app fiber.Router) {
	app.Post("/", h.createShortURL)
	app.Get("/:id", h.getFullURL)
}
