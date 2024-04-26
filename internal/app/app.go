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
			createShortUrl(w, r)
			return
		}

		http.Error(w, "", http.StatusBadRequest)
		return
	case http.MethodGet:
		if r.URL.Path == "/" {
			http.Error(w, "send id", http.StatusNotFound)
			return
		}

		getFullUrl(w, r)
		return
	default:
		http.Error(w, "method not allowed", http.StatusBadRequest)
		return
	}
}

func createShortUrl(w http.ResponseWriter, r *http.Request) {
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

	shortUrl := helper.CreateShortUrl(string(body))
	shortUrl = fmt.Sprintf("http://localhost:8080/%s", shortUrl)

	w.Header().Set("Content-Type", "text/html")
	w.WriteHeader(http.StatusCreated)

	_, err = w.Write([]byte(shortUrl))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	return
}

func getFullUrl(w http.ResponseWriter, r *http.Request) {
	url := r.URL.Path[1:]
	fullUrl := helper.GetShortUrl(url)

	if fullUrl == nil {
		http.Error(w, "short url not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Location", *fullUrl)
	w.WriteHeader(http.StatusTemporaryRedirect)

	return
}
