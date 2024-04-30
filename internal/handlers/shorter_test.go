package handlers

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"

	"github.com/vpbuyanov/short-url/internal/helper"
)

func TestCreateShortURL(t *testing.T) {
	var (
		log = logrus.New()
	)

	log.Level = logrus.DebugLevel

	tests := []struct {
		name   string
		path   string
		method string
		body   io.Reader
		assert func(w *http.Response)
	}{
		{
			name: "positive_test_#1",
			path: "/",
			body: io.MultiReader(
				bytes.NewReader([]byte("https://www.google.com")),
			),
			method: http.MethodPost,
			assert: func(w *http.Response) {
				assert.Equal(t, http.StatusCreated, w.StatusCode)
				assert.Equal(t, "text/plain", w.Header.Get("Content-Type"))
			},
		},
		{
			name:   "negative_test_#2",
			path:   "/",
			method: http.MethodPost,
			body:   nil,
			assert: func(w *http.Response) {
				defer func() {
					err := w.Body.Close()
					if err != nil {
						return
					}
				}()

				body, err := io.ReadAll(w.Body)
				if err != nil {
					return
				}

				assert.Equal(t, http.StatusBadRequest, w.StatusCode)
				assert.Equal(t, "text/plain; charset=utf-8", w.Header.Get("Content-Type"))
				assert.Equal(t, "empty body", string(body))
			},
		},
		{
			name:   "negative_test_#3",
			path:   "/",
			method: http.MethodGet,
			body:   nil,
			assert: func(w *http.Response) {
				defer func() {
					err := w.Body.Close()
					if err != nil {
						return
					}
				}()

				body, err := io.ReadAll(w.Body)
				if err != nil {
					return
				}

				assert.Equal(t, http.StatusMethodNotAllowed, w.StatusCode)
				assert.Equal(t, "text/plain; charset=utf-8", w.Header.Get("Content-Type"))
				assert.Equal(t, "Method Not Allowed", string(body))
			},
		},
	}

	a := fiber.New()
	a.Use(logger.New())

	url := helper.New()

	h := handlers{
		logger: log,
		url:    url,
	}

	a.Post("/", h.createShortURL)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(test.method, test.path, test.body)

			resp, err := a.Test(req, -1)
			if err != nil {
				return
			}

			test.assert(resp)
		})
	}

}

func TestGetFullURL(t *testing.T) {
	var (
		log = logrus.New()
	)

	log.Level = logrus.DebugLevel

	tests := []struct {
		name   string
		path   string
		method string
		header string
		assert func(w *http.Response)
	}{
		{
			name:   "positive_test_#1",
			path:   "/abcdefgG12",
			method: http.MethodGet,
			assert: func(w *http.Response) {
				assert.Equal(t, http.StatusTemporaryRedirect, w.StatusCode)
				assert.Equal(t, "https://google.com", w.Header.Get("Location"))
			},
		},
		{
			name:   "negative_test_#2",
			path:   "/as",
			method: http.MethodGet,
			assert: func(w *http.Response) {
				defer func() {
					err := w.Body.Close()
					if err != nil {
						return
					}
				}()

				body, err := io.ReadAll(w.Body)
				if err != nil {
					return
				}

				assert.Equal(t, http.StatusBadRequest, w.StatusCode)
				assert.Equal(t, "text/plain; charset=utf-8", w.Header.Get("Content-Type"))
				assert.Equal(t, "short url not found", string(body))
			},
		},
		{
			name:   "negative_test_#3",
			path:   "/",
			method: http.MethodGet,
			assert: func(w *http.Response) {
				defer func() {
					err := w.Body.Close()
					if err != nil {
						return
					}
				}()

				body, err := io.ReadAll(w.Body)
				if err != nil {
					return
				}

				assert.Equal(t, http.StatusNotFound, w.StatusCode)
				assert.Equal(t, "text/plain; charset=utf-8", w.Header.Get("Content-Type"))
				assert.Equal(t, "Cannot GET /", string(body))
			},
		},
		{
			name:   "negative_test_#4",
			path:   "/asd",
			method: http.MethodPost,
			assert: func(w *http.Response) {
				defer func() {
					err := w.Body.Close()
					if err != nil {
						return
					}
				}()

				body, err := io.ReadAll(w.Body)
				if err != nil {
					return
				}

				assert.Equal(t, http.StatusMethodNotAllowed, w.StatusCode)
				assert.Equal(t, "text/plain; charset=utf-8", w.Header.Get("Content-Type"))
				assert.Equal(t, "Method Not Allowed", string(body))
			},
		},
	}

	a := fiber.New()
	a.Use(logger.New())

	url := helper.New()

	h := handlers{
		logger: log,
		url:    url,
	}

	a.Get("/:id", h.getFullURL)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(test.method, test.path, nil)

			resp, err := a.Test(req, -1)
			if err != nil {
				return
			}

			test.assert(resp)
		})
	}
}
