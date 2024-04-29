package handlers

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
)

func TestCreateShortURL(t *testing.T) {
	var (
		logger = logrus.New()
	)

	logger.Level = logrus.DebugLevel

	tests := []struct {
		name   string
		path   string
		method string
		body   io.Reader
		assert func(w *httptest.ResponseRecorder)
	}{
		{
			name: "positive_test_#1",
			path: "/",
			body: io.MultiReader(
				bytes.NewReader([]byte("https://www.google.com")),
			),
			method: http.MethodPost,
			assert: func(w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusCreated, w.Code)
				assert.Equal(t, "text/plain", w.Header().Get("Content-Type"))
			},
		},
		{
			name:   "negative_test_#2",
			path:   "/",
			method: http.MethodPost,
			body:   nil,
			assert: func(w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, w.Code)
				assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))
				assert.Equal(t, "empty body\n", w.Body.String())
			},
		},
		{
			name:   "negative_test_#3",
			path:   "/",
			method: http.MethodGet,
			body:   nil,
			assert: func(w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
				assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))
				assert.Equal(t, "method not allowed\n", w.Body.String())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h := New(logger)

			request := httptest.NewRequest(test.method, test.path, test.body)

			w := httptest.NewRecorder()

			h.CreateShortURL(w, request)
			test.assert(w)
		})
	}
}

func TestGetFullURL(t *testing.T) {
	var (
		logger = logrus.New()
	)

	logger.Level = logrus.DebugLevel

	tests := []struct {
		name   string
		path   string
		method string
		header string
		assert func(w *httptest.ResponseRecorder)
	}{
		{
			name:   "positive_test_#1",
			path:   "/abcdefgG12",
			method: http.MethodGet,
			assert: func(w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusTemporaryRedirect, w.Code)
				assert.Equal(t, "https://google.com", w.Header().Get("Location"))
			},
		},
		{
			name:   "negative_test_#2",
			path:   "/as",
			method: http.MethodGet,
			assert: func(w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, w.Code)
				assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))
				assert.Equal(t, "short url not found\n", w.Body.String())
			},
		},
		{
			name:   "negative_test_#3",
			path:   "/",
			method: http.MethodGet,
			assert: func(w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, w.Code)
				assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))
				assert.Equal(t, "empty path url\n", w.Body.String())
			},
		},
		{
			name:   "negative_test_#4",
			path:   "/",
			method: http.MethodPost,
			assert: func(w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusMethodNotAllowed, w.Code)
				assert.Equal(t, "text/plain; charset=utf-8", w.Header().Get("Content-Type"))
				assert.Equal(t, "method not allowed\n", w.Body.String())
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			h := New(logger)

			request := httptest.NewRequest(test.method, test.path, nil)

			w := httptest.NewRecorder()
			h.GetFullURL(w, request)
			test.assert(w)
		})
	}
}
