package repos

type URL struct {
	shortURL map[string]string
	fullURL  map[string]string
}

func New() URL {
	shortURL := make(map[string]string)
	fullURL := make(map[string]string)

	return URL{
		shortURL: shortURL,
		fullURL:  fullURL,
	}
}

func (u *URL) GetShortURL(fullURL string) (string, bool) {
	shortURL, ok := u.fullURL[fullURL]

	return shortURL, ok
}

func (u *URL) SaveShortURL(fullURL, shortURL string) {
	u.shortURL[shortURL] = fullURL
	u.fullURL[fullURL] = shortURL
}

func (u *URL) GetFullURL(shortURL string) *string {
	url, ok := u.shortURL[shortURL]
	if ok {
		return &url
	}

	return nil
}
