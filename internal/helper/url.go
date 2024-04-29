package helper

import (
	"math/rand"
	"strconv"
	"strings"
)

type Url interface {
	CreateShortURL(url string) string
	GetShortURL(url string) *string
}

type url struct {
	urls map[string]string
}

func New() Url {
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
	code := make([]string, 10)
	for i := 0; i < 10; i++ {
		randNumber := rand.Intn(3)
		if randNumber == 1 {
			code[i] = strconv.Itoa(rand.Intn(10))
		} else if randNumber == 2 {
			code[i] = string(rune(rand.Intn(128)%26 + 65))
		} else {
			code[i] = string(rune(rand.Intn(128)%26 + 97))
		}
	}
	return strings.Join(code, "")
}
