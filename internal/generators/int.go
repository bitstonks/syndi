package generators

import (
	"fmt"
	"github.com/bitstonks/syndi/internal/config"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type IntUniformGenerator struct {
	rng      *rand.Rand
	nullable float64
	minVal   int // Inclusive
	maxVal   int // Non-inclusive
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
		rng:      rand.New(rand.NewSource(time.Now().UnixNano())),
		nullable: args.Nullable,
		minVal:   int(minVal),
		maxVal:   int(maxVal),
	}
}

func (g *IntUniformGenerator) Next() string {
	if g.nullable > 0 && g.rng.Float64() < g.nullable {
		return "NULL"
	}
	return fmt.Sprintf("%d", g.rng.Intn(g.maxVal-g.minVal)+g.minVal)
}
