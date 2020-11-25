package shortener_test

import (
	"testing"

	"github.com/pjonathas/url-shortener/shortener"
)

func TestShortener(t *testing.T) {

	t.Run("test Shorten.", func(t *testing.T) {
		storage := make(storageStub)
		shortenerService := shortener.Service{storage, generatorStub{}}

		url := "https://example.com"

		got, err := shortenerService.Shorten(url)

		want := "shorty"

		if got != want {
			t.Errorf("got %q want %q.", got, want)
		}

		if err != nil {
			t.Error("didn't expect an error.")
		}

		if shortenedURL, ok := storage[want]; !ok || shortenedURL.Original != url {
			t.Error("should have saved the url to storage.", url)
		}

	})

	t.Run("test Unshorten.", func(t *testing.T) {
		storage := make(storageStub)
		shortenerService := shortener.Service{storage, generatorStub{}}

		shortenerService.Shorten("https://example.com")

		got, err := shortenerService.Unshorten("shorty")
		want := "https://example.com"

		if err != nil {
			t.Error("didn't expect an error.")
		}

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	})

}
