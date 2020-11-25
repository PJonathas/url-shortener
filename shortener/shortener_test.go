package shortener_test

import (
	"testing"

	"github.com/pjonathas/url-shortener/shortener"
)

func TestShortener(t *testing.T) {
	shortenerService := shortener.Service{storageStub{}, generatorStub{}}

	got, err := shortenerService.Shorten("https://example.com")

	want := "shorty"

	if got != want {
		t.Errorf("got %q want %q.", got, want)
	}

	if err != nil {
		t.Error("didn't expect an error")
	}

}
