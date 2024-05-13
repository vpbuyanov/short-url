package handlers

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/stretchr/testify/assert"

	"github.com/vpbuyanov/short-url/internal/configs"
	"github.com/vpbuyanov/short-url/internal/repos"
	"github.com/vpbuyanov/short-url/internal/usecase"
)

func TestCreateShortURL(t *testing.T) {
	tests := []struct {
		name        string
		path        string
		method      string
		contentType string
		body        io.Reader
		assert      func(w *http.Response)
	}{
		{
			name: "positive_test_#1",
			path: "/",
			body: io.MultiReader(
				bytes.NewReader([]byte("https://www.google.com")),
			),
			contentType: "text/plain; charset=utf-8",
			method:      http.MethodPost,
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
			name:        "negative_test_#3",
			path:        "/",
			method:      http.MethodPost,
			body:        io.MultiReader(bytes.NewReader([]byte(` { "url" : "https://www.google.com" } `))),
			contentType: "application/json; charset=utf-8",
			assert: func(w *http.Response) {
				assert.Equal(t, http.StatusBadRequest, w.StatusCode)
			},
		},
		{
			name:   "negative_test_#4",
			path:   "/as",
			method: http.MethodPost,
			body: io.MultiReader(
				bytes.NewReader([]byte("https://www.google.com")),
			),
			contentType: "text/plain; charset=utf-8",
			assert: func(w *http.Response) {
				body, err := io.ReadAll(w.Body)
				if err != nil {
					return
				}

				assert.Equal(t, http.StatusNotFound, w.StatusCode)
				assert.Equal(t, "text/plain; charset=utf-8", w.Header.Get("Content-Type"))
				assert.Equal(t, "Not Found", string(body))
			},
		},
		{
			name:        "negative_test_#5",
			path:        "/",
			method:      http.MethodGet,
			body:        nil,
			contentType: "text/plain; charset=utf-8",
			assert: func(w *http.Response) {
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

	cfg := configs.Server{
		BaseURL: "http://localhost:8080",
		Address: "localhost:8080",
	}

	reposURL := repos.New()
	urlUC := usecase.New(reposURL, &cfg)

	h := Handlers{
		url: urlUC,
		cfg: &cfg,
	}

	a.Post("/", h.createShortURL)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(test.method, test.path, test.body)
			req.Header.Set("Content-Type", test.contentType)

			resp, err := a.Test(req, 10)
			if err != nil {
				return
			}

			test.assert(resp)

			defer func() {
				err = resp.Body.Close()
				if err != nil {
					return
				}
			}()
		})
	}
}

func TestGetFullURL(t *testing.T) {
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
				body, err := io.ReadAll(w.Body)
				if err != nil {
					return
				}

				assert.Equal(t, http.StatusBadRequest, w.StatusCode)
				assert.Equal(t, "text/plain; charset=utf-8", w.Header.Get("Content-Type"))
				assert.Equal(t, "not found url", string(body))
			},
		},
		{
			name:   "negative_test_#3",
			path:   "/",
			method: http.MethodGet,
			assert: func(w *http.Response) {
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
				body, err := io.ReadAll(w.Body)
				if err != nil {
					return
				}

				assert.Equal(t, http.StatusMethodNotAllowed, w.StatusCode)
				assert.Equal(t, "text/plain; charset=utf-8", w.Header.Get("Content-Type"))
				assert.Equal(t, "Method Not Allowed", string(body))
			},
		},
		{
			name:   "negative_test_#6",
			path:   "/1234567890102",
			method: http.MethodGet,
			assert: func(w *http.Response) {
				body, err := io.ReadAll(w.Body)
				if err != nil {
					return
				}

				assert.Equal(t, http.StatusBadRequest, w.StatusCode)
				assert.Equal(t, "text/plain; charset=utf-8", w.Header.Get("Content-Type"))
				assert.Equal(t, "not found url", string(body))
			},
		},
	}

	a := fiber.New()
	a.Use(logger.New())

	cfg := configs.Server{
		BaseURL: "http://localhost:8080",
		Address: "localhost:8080",
	}

	reposURL := repos.New()
	urlUC := usecase.New(reposURL, &cfg)

	reposURL.SaveShortURL("https://google.com", "abcdefgG12")

	h := Handlers{
		url: urlUC,
		cfg: &cfg,
	}

	a.Get("/:id", h.getFullURL)

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := httptest.NewRequest(test.method, test.path, http.NoBody)

			resp, err := a.Test(req, -1)
			if err != nil {
				return
			}

			test.assert(resp)

			defer func() {
				err = resp.Body.Close()
				if err != nil {
					return
				}
			}()
		})
	}
}
