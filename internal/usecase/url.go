package usecase

import (
	"errors"
	"fmt"
	"math/rand"
	neturl "net/url"

	"github.com/vpbuyanov/short-url/internal/configs"
	"github.com/vpbuyanov/short-url/internal/repos"
)

const (
	lenHash = 10
)

var strShortURL = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type URL struct {
	urlRepos repos.URL
	cfg      *configs.Server
}

func New(urlRepos repos.URL, cfg *configs.Server) *URL {
	return &URL{
		urlRepos: urlRepos,
		cfg:      cfg,
	}
}

func (u *URL) CreateAndSaveShortURL(url string) (*string, error) {
	shortURL, ok := u.urlRepos.GetShortURL(url)
	if ok {
		return &shortURL, nil
	}

	for range 3 {
		shortURL = u.generateShortURL()

		fullURL := u.urlRepos.GetFullURL(shortURL)
		if fullURL != nil {
			continue
		}

		u.urlRepos.SaveShortURL(url, shortURL)

		path, err := neturl.JoinPath(u.cfg.BaseURL, shortURL)
		if err != nil {
			return nil, fmt.Errorf("can't join path, err: %w", err)
		}

		return &path, nil
	}

	return nil, errors.New("can not generate unique short URL")
}

func (u *URL) GetFullURL(url string) (*string, error) {
	fullURL := u.urlRepos.GetFullURL(url)

	if fullURL == nil {
		return nil, errors.New("not found url")
	}

	return fullURL, nil
}

func (u *URL) generateShortURL() string {
	b := make([]rune, lenHash)

	for i := range b {
		b[i] = strShortURL[rand.Intn(len(strShortURL))]
	}

	return string(b)
}
