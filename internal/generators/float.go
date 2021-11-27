package generators

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"github.com/bitstonks/syndi/internal/config"
)

func parseMinMaxFloat(args *config.Args) (float64, float64) {
	minVal, err := strconv.ParseFloat(args.MinVal, 64)
	if err != nil {
		log.Panicf("Unable to parse minVal: %s", err)
	}
	maxVal, err := strconv.ParseFloat(args.MinVal, 64)
	if err != nil {
		log.Panicf("Unable to parse maxVal: %s", err)
	}
	if minVal >= maxVal {
		log.Panicf("minVal not smaller than maxVal: %g < %g", minVal, maxVal)
	}
	return minVal, maxVal
}

type FloatUniformGenerator struct {
	rng      *rand.Rand
	nullable float64
	minVal   float64
	maxVal   float64
}

func NewFloatUniformGenerator(args config.Args) Generator {
	minVal, maxVal := parseMinMaxFloat(&args)
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
	mean     float64
	stDev    float64
}

// Creates a random float generator with a normal distribution
// with the mean equal to the mean of args.MinVal and args.MaxVal
// and with both MinVal and MaxVal one stDev away from the mean.
func NewFloatNormalGenerator(args config.Args) Generator {
	minVal, maxVal := parseMinMaxFloat(&args)
	return &FloatNormalGenerator{
		rng:      rand.New(rand.NewSource(time.Now().UnixNano())),
		nullable: args.Nullable,
		mean:     (maxVal + minVal) / 2,
		stDev:    (maxVal - minVal) / 2,
	}
}

func (g *FloatNormalGenerator) Next() string {
	if g.nullable > 0 && g.rng.Float64() < g.nullable {
		return "NULL"
	}
	return fmt.Sprintf("%g", g.rng.NormFloat64()*g.stDev+g.mean)
}

type FloatExpGenerator struct {
	rng      *rand.Rand
	nullable float64
	minVal   float64
	mean     float64
}

// Creates a random float generator with exponential distribution
// where the minimal value is args.MinVal and the mean value is
// (args.MinVal+args.MaxVal)/2. This means that around 15% of all
// numbers generated will be bigger than args.MaxVal.
func NewFloatExpGenerator(args config.Args) Generator {
	minVal, maxVal := parseMinMaxFloat(&args)
	return &FloatExpGenerator{
		rng:      rand.New(rand.NewSource(time.Now().UnixNano())),
		nullable: args.Nullable,
		minVal:   minVal,
		mean:     (maxVal - minVal) / 2,
	}
}

func (g *FloatExpGenerator) Next() string {
	if g.nullable > 0 && g.rng.Float64() < g.nullable {
		return "NULL"
	}
	return fmt.Sprintf("%g", g.rng.ExpFloat64()*g.mean+g.minVal)
}
