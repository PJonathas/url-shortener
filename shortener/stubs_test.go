package shortener_test

import (
	"errors"

	"github.com/pjonathas/url-shortener/shortener"
)

type generatorStub struct{}

func (g generatorStub) ID() (string, error) {
	return "shorty", nil
}

type storageStub map[string]shortener.URL

func (s storageStub) Put(shortenerURL shortener.URL) error {
	s[shortenerURL.ID] = shortenerURL

	return nil
}

func (s storageStub) Get(code string) (shortener.URL, error) {

	shortenerURL, ok := s[code]

	if !ok {
		return shortener.URL{}, errors.New("url not found")
	}

	return shortenerURL, nil
}
