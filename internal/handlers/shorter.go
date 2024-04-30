package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

func (h *handlers) createShortURL(ctx *fiber.Ctx) error {
	body := ctx.BodyRaw()

	if len(body) == 0 {
		h.logger.Error("empty body")
		return ctx.Status(fiber.StatusBadRequest).SendString("empty body")
	}

	shortURL := h.url.CreateShortURL(string(body))
	shortURL = fmt.Sprintf("http://localhost:8080/%s", shortURL)

	ctx.Set("Content-Type", "text/plain")

	return ctx.Status(fiber.StatusCreated).SendString(shortURL)
}

func (h *handlers) getFullURL(ctx *fiber.Ctx) error {
	urls := ctx.AllParams()
	url := urls["id"]
	if len(url) == 0 {
		h.logger.Error("empty url")
		return ctx.Status(fiber.StatusBadRequest).SendString("empty url")
	}

	fullURL := h.url.GetShortURL(url)

	if fullURL == nil {
		h.logger.Error("short url not found")
		return ctx.Status(http.StatusBadRequest).SendString("short url not found")
	}

	ctx.Status(http.StatusTemporaryRedirect).Set("Location", *fullURL)
	return ctx.Send(nil)
}
