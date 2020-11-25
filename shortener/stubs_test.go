package shortener_test

import "github.com/pjonathas/url-shortener/shortener"

type generatorStub struct{}

func (g generatorStub) Code() (string, error) {
	return "shorty", nil
}

type storageStub map[string]shortener.URL

func (s storageStub) Put(shortener.URL) error {

	return nil
}

func (s storageStub) Get(code string) (shortener.URL, error) {

	return shortener.URL{}, nil
}
