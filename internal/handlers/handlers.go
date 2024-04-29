package handlers

import (
	"github.com/vpbuyanov/short-url/internal/helper"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Handler interface {
	Shorter(w http.ResponseWriter, r *http.Request)
	CreateShortURL(w http.ResponseWriter, r *http.Request)
	GetFullURL(w http.ResponseWriter, r *http.Request)
}

type handlers struct {
	logger *logrus.Logger
	url    helper.Url
}

func New(log *logrus.Logger) Handler {
	url := helper.New()

	return &handlers{
		logger: log,
		url:    url,
	}
}
