package handlers

import (
	"fmt"
	"net/http"

	"github.com/gofiber/fiber/v2"
)

const (
	canNotSendRequest = "can not send request, err: %w"
)

func (h *Handlers) createShortURL(ctx *fiber.Ctx) error {
	body := ctx.BodyRaw()

	if len(body) == 0 {
		err := ctx.Status(fiber.StatusBadRequest).SendString("empty body")
		if err != nil {
			return fmt.Errorf(canNotSendRequest, err)
		}

		return nil
	}

	url, err := h.url.CreateAndSaveShortURL(string(body))
	if err != nil {
		return fmt.Errorf("can not create and save url, err: %w", err)
	}

	ctx.Set("Content-Type", "text/plain")

	err = ctx.Status(fiber.StatusCreated).SendString(*url)
	if err != nil {
		return fmt.Errorf(canNotSendRequest, err)
	}

	return nil
}

func (h *Handlers) getFullURL(ctx *fiber.Ctx) error {
	urls := ctx.AllParams()
	getURL := urls["id"]
	if len(getURL) == 0 {
		err := ctx.Status(fiber.StatusBadRequest).SendString("empty url")
		if err != nil {
			return fmt.Errorf(canNotSendRequest, err)
		}

		return nil
	}

	fullURL, err := h.url.GetFullURL(getURL)
	if err != nil {
		err = ctx.Status(http.StatusBadRequest).SendString(err.Error())
		if err != nil {
			return fmt.Errorf(canNotSendRequest, err)
		}

		return nil
	}

	ctx.Status(http.StatusTemporaryRedirect).Set("Location", *fullURL)
	err = ctx.Send(nil)
	if err != nil {
		return fmt.Errorf(canNotSendRequest, err)
	}

	return nil
}
