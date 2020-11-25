package shortener

// Service takes our dependencies
type Service struct {
	Storage   Storage
	Generator Generator
}

// URL takes a url and returns a short code
func (s Service) Shorten(originalUrl string) (string, error) {
	shortened, err := s.Generator.Code()
	if err != nil {
		return "", nil
	}

	url := URL{originalUrl, shortened}

	return url.Shortened, s.Storage.Put(url)

}

func (s Service) Unshorten(code string) (string, error) {

	return "", nil
}
