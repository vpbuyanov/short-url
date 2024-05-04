package services

import (
	"math/rand"
	"strconv"
	"strings"
)

const (
	lenHash            = 10
	numberOfComponents = 3
	numbers            = 10
	uppercase          = 65
	lowercase          = 97
	ascii              = 128
)

type URL interface {
	CreateShortURL(url string) string
	GetShortURL(url string) *string
}

type url struct {
	urls map[string]string
}

func New() URL {
	newUrls := make(map[string]string)
	newUrls["abcdefgG12"] = "https://google.com"

	return &url{
		urls: newUrls,
	}
}

func (u *url) CreateShortURL(url string) string {
	var res string

	for {
		res = u.generateURL()

		_, ok := u.urls[res]
		if !ok {
			break
		}
	}

	u.urls[res] = url

	return res
}

func (u *url) GetShortURL(url string) *string {
	getURL, ok := u.urls[url]
	if ok {
		return &getURL
	}

	return nil
}

func (u *url) generateURL() string {
	code := make([]string, lenHash)
	for i := range lenHash {
		randNumber := rand.Intn(numberOfComponents)

		switch randNumber {
		case 0:
			code[i] = strconv.Itoa(rand.Intn(numbers))
		case 1:
			code[i] = string(rune(rand.Intn(ascii)%26 + uppercase))
		default:
			code[i] = string(rune(rand.Intn(ascii)%26 + lowercase))
		}
	}
	return strings.Join(code, "")
}
