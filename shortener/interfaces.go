package shortener

type Storage interface {
	Put(URL) error
	Get(string) (URL, error)
}

type Generator interface {
	Code() (string, error)
}
