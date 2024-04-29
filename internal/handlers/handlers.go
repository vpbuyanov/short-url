package handlers

import (
	"net/http"

	"github.com/sirupsen/logrus"

	"github.com/vpbuyanov/short-url/internal/helper"
)

type Handler interface {
	Shorter(w http.ResponseWriter, r *http.Request)
	CreateShortURL(w http.ResponseWriter, r *http.Request)
	GetFullURL(w http.ResponseWriter, r *http.Request)
}

type handlers struct {
	logger *logrus.Logger
	url    helper.URL
}

func New(log *logrus.Logger) Handler {
	url := helper.New()

	return &handlers{
		logger: log,
		url:    url,
	}
}
