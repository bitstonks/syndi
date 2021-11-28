package generators

import (
	"math/rand"

	"github.com/bitstonks/syndi/internal/config"
)

// yaay, globals!
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type StringGenerator struct {
	rng *rand.Rand
	len int
}

func NewStringGenerator(args config.Args) Generator {
	// TODO: when defined use args.OneOf for character selection instead of global letters
	return &StringGenerator{
		rng: NewRng(),
		len: args.Length,
	}
}

func (g *StringGenerator) Next() string {
	b := make([]rune, g.len)
	for i := range b {
		b[i] = letters[g.rng.Intn(len(letters))]
	}
	return string(b)
}
