package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const (
	canNotSendRequest = "can not send request, err: %w"
)

func (h *handlers) createShortURL(ctx *fiber.Ctx) error {
	body := ctx.BodyRaw()

	if len(body) == 0 {
		h.logger.Error("empty body")

		err := ctx.Status(fiber.StatusBadRequest).SendString("empty body")
		if err != nil {
			return fmt.Errorf(canNotSendRequest, err)
		}

		return nil
	}

	shortURL := h.url.CreateShortURL(string(body))
	shortURL = h.cfg.ResShortURL + "/" + shortURL

	ctx.Set("Content-Type", "text/plain")

	err := ctx.Status(fiber.StatusCreated).SendString(shortURL)
	if err != nil {
		return fmt.Errorf(canNotSendRequest, err)
	}

	return nil
}

func (h *handlers) getFullURL(ctx *fiber.Ctx) error {
	urls := ctx.AllParams()
	url := urls["id"]
	if len(url) == 0 {
		h.logger.Error("empty url")
		err := ctx.Status(fiber.StatusBadRequest).SendString("empty url")
		if err != nil {
			return fmt.Errorf(canNotSendRequest, err)
		}

		return nil
	}

	fullURL := h.url.GetShortURL(url)

	if fullURL == nil {
		h.logger.Error("short url not found")
		err := ctx.Status(http.StatusBadRequest).SendString("short url not found")
		if err != nil {
			return fmt.Errorf(canNotSendRequest, err)
		}

		return nil
	}

	ctx.Status(http.StatusTemporaryRedirect).Set("Location", *fullURL)
	err := ctx.Send(nil)
	if err != nil {
		return fmt.Errorf(canNotSendRequest, err)
	}

	return nil
}
