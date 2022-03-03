package generators

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/bitstonks/syndi/internal/config"
)

type intUniformGenerator struct {
	rng    *rand.Rand
	minVal int // Inclusive
	spread int // minVal + spread non-inclusive
}

func NewIntUniformGenerator(args config.ColumnDef) Generator {
	minVal, err := strconv.ParseInt(args.MinVal, 10, 64)
	if err != nil {
		log.Panicf("Unable to parse minVal: %s", err)
	}
	maxVal, err := strconv.ParseInt(args.MaxVal, 10, 64)
	if err != nil {
		log.Panicf("Unable to parse maxVal: %s", err)
	}
	if minVal >= maxVal {
		log.Panicf("minVal not smaller than maxVal: %d < %d", minVal, maxVal)
	}
	return &intUniformGenerator{
		rng:    newRng(),
		minVal: int(minVal),
		spread: int(maxVal - minVal),
	}
}

func (g *intUniformGenerator) Next() interface{} {
	return g.rng.Intn(g.spread) + g.minVal
}
