package shortener

// Storage is a dependency to our shortener service
type Storage interface {
	Put(URL) error
	Get(string) (URL, error)
}

// Generator is a dependency to our shortener service
type Generator interface {
	Code() (string, error)
}
