package generators

import (
	"log"
	"math/rand"
	"strconv"
	"strings"

	"github.com/bitstonks/syndi/internal/config"
)

type OneOfGenerator struct {
	rng     *rand.Rand
	weights map[string]int
	total   int
}

func NewOneOfGenerator(args config.Args) Generator {
	weights := getMultipleChoice(args.OneOf)
	total := 0
	for _, w := range weights {
		total += w
	}
	return &OneOfGenerator{
		rng:     NewRng(),
		weights: weights,
		total:   total,
	}
}

func (g *OneOfGenerator) Next() string {
	n := g.rng.Intn(g.total)
	for v, w := range g.weights {
		n -= w
		if n < 0 {
			return v
		}
	}
	return ""
}

func getMultipleChoice(opts string) (weights map[string]int) {
	for _, opt := range strings.Split(opts, ";") {
		parts := strings.Split(opt, ":")
		w, err := strconv.ParseInt(parts[len(parts)-1], 10, 64)
		if len(parts) == 1 || err != nil {
			if len(parts) != 1 && err != nil {
				log.Printf("unable to parse option %s defaulting to weight 1: %s", opt, err)
			}
			weights[opt] = 1
			continue
		}
		if w <= 0 {
			log.Panicf("weight should be %q > 0: %s", w, opt)
		}
		weights[strings.Join(parts[:len(parts)-1], ":")] = int(w)
	}
	if len(weights) == 0 {
		log.Panicf("unable to parse even a single option for multiple choice type")
	}
	return
}
