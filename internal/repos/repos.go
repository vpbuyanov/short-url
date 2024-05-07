package repos

import "github.com/vpbuyanov/short-url/internal/services"

type URL struct {
	url     services.URL
	saveURL map[string]string
}

func New(service services.URL) URL {
	saveURL := make(map[string]string)
	saveURL["abcdefgG12"] = "https://google.com"

	return URL{
		url:     service,
		saveURL: saveURL,
	}
}

func (u *URL) CreateShortURL(url string) string {
	var res string

	for range 3 {
		res = u.url.GenerateShortURL()

		_, ok := u.saveURL[res]
		if !ok {
			break
		}
	}

	u.saveURL[res] = url

	return res
}

func (u *URL) GetShortURL(url string) *string {
	getURL, ok := u.saveURL[url]
	if ok {
		return &getURL
	}

	return nil
}
