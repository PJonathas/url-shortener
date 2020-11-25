package shortener

// URL is the type to link the Original url with the Shortened version
type URL struct {
	Original string `bson:"url"`
	ID       string `bson:"_id"`
}
