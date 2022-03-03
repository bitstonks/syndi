package generators

import (
	"math/rand"
)

type nullifier struct {
	rng      *rand.Rand
	nullable float64
	gen      Generator
}

func MakeNullifier(gen Generator, nullable float64) Generator {
	if nullable <= 0 {
		return gen
	}
	return &nullifier{
		rng:      newRng(),
		nullable: nullable,
		gen:      gen,
	}
}

func (n *nullifier) Next() interface{} {
	if n.rng.Float64() < n.nullable {
		return "NULL"
	}
	return n.gen.Next()
}
