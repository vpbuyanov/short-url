package handlers

import (
	"net/http"

	"github.com/sirupsen/logrus"
)

type Handler interface {
	Shorter(w http.ResponseWriter, r *http.Request)
}

type handlers struct {
	logger *logrus.Logger
}

func New(log *logrus.Logger) Handler {
	return &handlers{
		logger: log,
	}
}
