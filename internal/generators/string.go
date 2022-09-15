package generators

import (
	"math/rand"

	"github.com/bitstonks/syndi/internal/config"
)

type stringGenerator struct {
	rng *rand.Rand
	len int
	all []rune
	quotedFmt
}

func NewStringGenerator(args config.ColumnDef) Generator {
	all := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	if len(args.OneOf) > 0 {
		all = []rune(args.OneOf)
	}
	return &stringGenerator{
		rng: newRng(),
		len: args.Length,
		all: all,
	}
}

func (g *stringGenerator) Next() interface{} {
	b := make([]rune, g.len)
	for i := range b {
		b[i] = g.all[g.rng.Intn(len(g.all))]
	}
	return string(b)

}
