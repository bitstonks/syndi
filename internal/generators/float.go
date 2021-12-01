package generators

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"

	"github.com/bitstonks/syndi/internal/config"
)

func parseMinMaxFloat(args *config.Args) (float64, float64) {
	minVal, err := strconv.ParseFloat(args.MinVal, 64)
	if err != nil {
		log.Panicf("Unable to parse minVal: %s", err)
	}
	maxVal, err := strconv.ParseFloat(args.MaxVal, 64)
	if err != nil {
		log.Panicf("Unable to parse maxVal: %s", err)
	}
	if minVal >= maxVal {
		log.Panicf("minVal not smaller than maxVal: %g < %g", minVal, maxVal)
	}
	return minVal, maxVal
}

type floatUniformGenerator struct {
	rng    *rand.Rand
	minVal float64
	spread float64
}

func NewFloatUniformGenerator(args config.Args) Generator {
	minVal, maxVal := parseMinMaxFloat(&args)
	return &floatUniformGenerator{
		rng:    newRng(),
		minVal: minVal,
		spread: maxVal - minVal,
	}
}

func (g *floatUniformGenerator) Next() string {
	v := g.minVal + g.rng.Float64()*g.spread
	return fmt.Sprintf("%g", v)
}

type floatNormalGenerator struct {
	rng   *rand.Rand
	mean  float64
	stDev float64
}

// Creates a random float generator with a normal distribution
// with the mean equal to the mean of args.MinVal and args.MaxVal
// and with both MinVal and MaxVal one stDev away from the mean.
func NewFloatNormalGenerator(args config.Args) Generator {
	minVal, maxVal := parseMinMaxFloat(&args)
	return &floatNormalGenerator{
		rng:   newRng(),
		mean:  (maxVal + minVal) / 2,
		stDev: (maxVal - minVal) / 2,
	}
}

func (g *floatNormalGenerator) Next() string {
	return fmt.Sprintf("%g", g.rng.NormFloat64()*g.stDev+g.mean)
}

type floatExpGenerator struct {
	rng    *rand.Rand
	minVal float64
	mean   float64
}

// Creates a random float generator with exponential distribution
// where the minimal value is args.MinVal and the mean value is
// (args.MinVal+args.MaxVal)/2. This means that around 15% of all
// numbers generated will be bigger than args.MaxVal.
func NewFloatExpGenerator(args config.Args) Generator {
	minVal, maxVal := parseMinMaxFloat(&args)
	return &floatExpGenerator{
		rng:    newRng(),
		minVal: minVal,
		mean:   (maxVal - minVal) / 2,
	}
}

func (g *floatExpGenerator) Next() string {
	return fmt.Sprintf("%g", g.rng.ExpFloat64()*g.mean+g.minVal)
}
