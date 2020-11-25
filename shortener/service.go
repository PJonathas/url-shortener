package shortener

// Service takes our dependencies
type Service struct {
	Storage   Storage
	Generator Generator
}

// Shorten takes a url string and returns a short code
func (s Service) Shorten(originalURL string) (string, error) {
	id, err := s.Generator.ID()
	if err != nil {
		return "", nil
	}

	url := URL{originalURL, id}

	return url.ID, s.Storage.Put(url)

}

// Unshorten takes short code and returns a url
func (s Service) Unshorten(id string) (string, error) {
	url, err := s.Storage.Get(id)
	if err != nil {
		return "", err
	}

	return url.Original, nil

}
