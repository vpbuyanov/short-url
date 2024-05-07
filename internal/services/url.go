package services

import (
	"math/rand"
	"strings"
)

const (
	lenHash = 10
)

type URL struct{}

func New() URL {
	return URL{}
}

func (u *URL) GenerateShortURL() string {
	strShortURL := []rune("abcdefghijklmnopqrstuvwxyz" + "ABCDEFGHIJKLMNOPQRSTUVWXYZ" + "0123456789")

	var b strings.Builder

	for range lenHash {
		b.WriteRune(strShortURL[rand.Intn(len(strShortURL))])
	}

	return b.String()
}
