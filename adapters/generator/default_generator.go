package generator

import "github.com/aidarkhanov/nanoid/v2"

const (
	validStrings = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
)

// Default is the concrete type the satisfies
// the generator interface
type Default struct{}

// GenerateCode generates a new code
func (g Default) ID() (string, error) {

	return nanoid.GenerateString(validStrings, 8)
}
