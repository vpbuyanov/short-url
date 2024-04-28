package handlers

import (
	"fmt"
	"io"
	"net/http"

	"github.com/vpbuyanov/short-url/internal/helper"
)

func (h *handlers) Shorter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p := r.URL.Path

		if p == "/" {
			h.createShortURL(w, r)
			h.logger.Infof("created short url")
			return
		}

		h.logger.Error("status bad request, page not found")
		http.Error(w, "status bad request, page not found", http.StatusBadRequest)
		return
	case http.MethodGet:
		if r.URL.Path == "/" {
			h.logger.Error("StatusNotFound, not send id")
			http.Error(w, "send id", http.StatusNotFound)
			return
		}

		h.getFullURL(w, r)
		h.logger.Info("sent short url")
		return
	default:
		h.logger.Error("unknown method, status bad request")
		http.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}
}

func (h *handlers) createShortURL(w http.ResponseWriter, r *http.Request) {
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}(r.Body)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(body) == 0 {
		http.Error(w, "empty body", http.StatusBadRequest)
		return
	}

	shortURL := helper.CreateShortURL(string(body))
	shortURL = fmt.Sprintf("http://localhost:8080/%s", shortURL)

	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)

	_, err = w.Write([]byte(shortURL))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func (h *handlers) getFullURL(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path[1:]
	fullURL := helper.GetShortURL(url)

	if fullURL == nil {
		http.Error(w, "short url not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Location", *fullURL)
	w.WriteHeader(http.StatusTemporaryRedirect)
}
