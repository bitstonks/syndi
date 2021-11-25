package generators

import (
	"fmt"
	"github.com/bitstonks/syndi/internal/config"
	"log"
	"math/rand"
	"strconv"
	"time"
)

type FloatUniformGenerator struct {
	rng      *rand.Rand
	nullable float64
	minVal   float64
	maxVal   float64
}

func NewFloatUniformGenerator(args config.Args) Generator {
	minVal, err := strconv.ParseFloat(args.MinVal, 64)
	if err != nil {
		log.Panicf("Unable to parse minVal: %s", err)
	}
	maxVal, err := strconv.ParseFloat(args.MinVal, 64)
	if err != nil {
		log.Panicf("Unable to parse minVal: %s", err)
	}
	if minVal >= maxVal {
		log.Panicf("minVal not smaller than maxVal: %g < %g", minVal, maxVal)
	}
	return &FloatUniformGenerator{
		rng:      rand.New(rand.NewSource(time.Now().UnixNano())),
		nullable: args.Nullable,
		minVal:   minVal,
		maxVal:   maxVal,
	}
}

func (g *FloatUniformGenerator) Next() string {
	if g.nullable > 0 && g.rng.Float64() < g.nullable {
		return "NULL"
	}
	v := g.minVal + g.rng.Float64()*(g.maxVal-g.minVal)
	return fmt.Sprintf("%g", v)
}

type FloatNormalGenerator struct {
	rng      *rand.Rand
	nullable float64
}

func NewFloatNormalGenerator(args config.Args) Generator {
	return &FloatNormalGenerator{
		rng:      rand.New(rand.NewSource(time.Now().UnixNano())),
		nullable: args.Nullable,
	}
}

func (g *FloatNormalGenerator) Next() string {
	if g.nullable > 0 && g.rng.Float64() < g.nullable {
		return "NULL"
	}
	return fmt.Sprintf("%g", g.rng.NormFloat64())
}
