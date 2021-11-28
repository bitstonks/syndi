package generators

import (
	"math/rand"
)

type Nullable struct {
	rng      *rand.Rand
	nullable float64
	gen      Generator
}

func MakeNullable(gen Generator, nullable float64) Generator {
	if nullable <= 0 {
		return gen
	}
	return &Nullable{
		rng:      NewRng(),
		nullable: nullable,
		gen:      gen,
	}
}

func (n *Nullable) Next() string {
	if n.rng.Float64() < n.nullable {
		return "NULL"
	}
	return n.gen.Next()
}
