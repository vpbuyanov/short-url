package handlers

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/gofiber/fiber/v2"
)

const (
	canNotSendRequest = "can not send request, err: %w"
)

func (h *Handlers) createShortURL(ctx *fiber.Ctx) error {
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
	path, err := url.JoinPath(h.cfg.BaseURL, shortURL)
	if err != nil {
		return fmt.Errorf("can't join path, err: %w", err)
	}

	ctx.Set("Content-Type", "text/plain")

	err = ctx.Status(fiber.StatusCreated).SendString(path)
	if err != nil {
		return fmt.Errorf(canNotSendRequest, err)
	}

	return nil
}

func (h *Handlers) getFullURL(ctx *fiber.Ctx) error {
	urls := ctx.AllParams()
	getURL := urls["id"]
	if len(getURL) == 0 {
		h.logger.Error("empty url")
		err := ctx.Status(fiber.StatusBadRequest).SendString("empty url")
		if err != nil {
			return fmt.Errorf(canNotSendRequest, err)
		}

		return nil
	}

	fullURL := h.url.GetShortURL(getURL)

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
