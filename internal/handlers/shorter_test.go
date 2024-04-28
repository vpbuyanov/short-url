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
	type fields struct {
		logger *logrus.Logger
	}

	var (
		logger = logrus.New()
	)

	logger.Level = logrus.DebugLevel

	tests := []struct {
		name   string
		fields fields
		path   string
		method string
		body   io.Reader
		assert func(w *httptest.ResponseRecorder)
	}{
		{
			name: "positive test #1",
			fields: fields{
				logger: logger,
			},
			path: "/",
			body: io.MultiReader(
				bytes.NewReader([]byte("...my header...")),
			),
			method: http.MethodPost,
			assert: func(w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusCreated, w.Code)
			},
		},
		{
			name: "negative_test_#2",
			fields: fields{
				logger: logger,
			},
			path: "/shortest",
			assert: func(w *httptest.ResponseRecorder) {
				assert.Equal(t, http.StatusBadRequest, w.Code)
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()

			h := &handlers{
				logger: test.fields.logger,
			}

			request := httptest.NewRequest(test.method, test.path, test.body)

			w := httptest.NewRecorder()
			h.createShortURL(w, request)
			test.assert(w)
		})
	}
}

func TestGetFullURL(t *testing.T) {

}
