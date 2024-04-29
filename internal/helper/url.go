package helper

import (
	"math/rand"
	"strconv"
	"strings"
)

var (
	urls map[string]string
)

func init() {
	urls = make(map[string]string)
	urls["abcdefgG12"] = "https://google.com"
}

func CreateShortURL(url string) string {
	var res string

	for {
		res = generateURL()

		_, ok := urls[res]
		if !ok {
			break
		}
	}

	urls[res] = url

	return res
}

func GetShortURL(url string) *string {
	getURL, ok := urls[url]
	if ok {
		return &getURL
	}

	return nil
}

func generateURL() string {
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
