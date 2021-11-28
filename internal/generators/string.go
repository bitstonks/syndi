package generators

import (
	"math/rand"

	"github.com/bitstonks/syndi/internal/config"
)

type StringGenerator struct {
	rng *rand.Rand
	len int
	all []rune
}

func NewStringGenerator(args config.Args) Generator {
	all := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	if len(args.OneOf) > 0 {
		all = []rune(args.OneOf)
	}
	return &StringGenerator{
		rng: NewRng(),
		len: args.Length,
		all: all,
	}
}

func (g *StringGenerator) Next() string {
	b := make([]rune, g.len)
	for i := range b {
		b[i] = g.all[g.rng.Intn(len(g.all))]
	}
	return string(b)
}
