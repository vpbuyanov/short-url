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

type URL struct{}

func New() URL {
	return URL{}
}

func (u *URL) GenerateShortURL() string {
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
