package generators

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/bitstonks/syndi/internal/config"
)

type IntUniformGenerator struct {
	rng    *rand.Rand
	minVal int // Inclusive
	spread int // minVal + spread non-inclusive
}

func NewIntUniformGenerator(args config.Args) Generator {
	minVal, err := strconv.ParseInt(args.MinVal, 10, 64)
	if err != nil {
		log.Panicf("Unable to parse minVal: %s", err)
	}
	maxVal, err := strconv.ParseInt(args.MinVal, 10, 64)
	if err != nil {
		log.Panicf("Unable to parse minVal: %s", err)
	}
	if minVal >= maxVal {
		log.Panicf("minVal not smaller than maxVal: %d < %d", minVal, maxVal)
	}
	return &IntUniformGenerator{
		rng:    NewRng(),
		minVal: int(minVal),
		spread: int(maxVal - minVal),
	}
}

func (g *IntUniformGenerator) Next() string {
	return fmt.Sprintf("%d", g.rng.Intn(g.spread)+g.minVal)
}
