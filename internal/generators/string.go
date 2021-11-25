package generators

import (
	"github.com/bitstonks/syndi/internal/config"
	"math/rand"
	"time"
)

// yaay, globals!
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

type StringGenerator struct {
	rng      *rand.Rand
	len      int
	nullable float64
}

func NewStringGenerator(args config.Args) Generator {
	// TODO: when defined use args.OneOf for character selection instead of global letters
	return &StringGenerator{
		rng:      rand.New(rand.NewSource(time.Now().UnixNano())),
		nullable: args.Nullable,
		len:      args.Length,
	}
}

func (g *StringGenerator) Next() string {
	if g.nullable > 0 && g.rng.Float64() < g.nullable {
		return "NULL"
	}
	b := make([]rune, g.len)
	for i := range b {
		b[i] = letters[g.rng.Intn(len(letters))]
	}
	return string(b)
}
