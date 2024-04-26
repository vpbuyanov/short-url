package app

import (
	"fmt"
	"io"
	"net/http"

	"github.com/vpbuyanov/short-url/internal/helper"
)

type App struct{}

func New() *App {
	return new(App)
}

func (app *App) Start() error {
	http.HandleFunc("/", handlerShorter)

	return http.ListenAndServe(":8080", nil)
}

func handlerShorter(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		p := r.URL.Path

		if p == "/" {
			createShortURL(w, r)
			return
		}

		http.Error(w, "", http.StatusBadRequest)
		return
	case http.MethodGet:
		if r.URL.Path == "/" {
			http.Error(w, "send id", http.StatusNotFound)
			return
		}

		getFullURL(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}
}

func createShortURL(w http.ResponseWriter, r *http.Request) {
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

	shortURL := helper.CreateShortURL(string(body))
	shortURL = fmt.Sprintf("http://localhost:8080/%s", shortURL)

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusCreated)

	_, err = w.Write([]byte(shortURL))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	return
}

func getFullURL(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path[1:]
	fullURL := helper.GetShortURL(url)

	if fullURL == nil {
		http.Error(w, "short url not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Location", *fullURL)
	w.WriteHeader(http.StatusTemporaryRedirect)

	return
}
