package generators

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"

	"github.com/bitstonks/syndi/internal/config"
)

// OneOfGenerator generates values from a finite pool of predefined choices.
// Choices can be selected uniformly or weighted to select some more often than others.
type OneOfGenerator struct {
	rng     *rand.Rand
	weights []weighted
	total   int
}

// NewOneOfGenerator constructs a OneOfGenerator
func NewOneOfGenerator(args config.Args) Generator {
	weights, total := getMultipleChoice(args.OneOf)
	return &OneOfGenerator{
		rng:     NewRng(),
		weights: weights,
		total:   total,
	}
}

func (g *OneOfGenerator) Next() string {
	n := g.rng.Intn(g.total)
	for _, w := range g.weights {
		n -= w.weight
		if n < 0 {
			return w.name
		}
	}
	return ""
}

// QuotedOneOfGenerator is a light wrapper for OneOfGenerator that wraps the generated strings in single quotes.
type QuotedOneOfGenerator struct {
	gen Generator
}

func NewQuotedOneOfGenerator(args config.Args) Generator {
	return &QuotedOneOfGenerator{
		gen: NewOneOfGenerator(args),
	}
}

func (g *QuotedOneOfGenerator) Next() string {
	return fmt.Sprintf("'%s'", g.gen.Next())
}

type weighted struct {
	name   string
	weight int
}

func getMultipleChoice(opts string) (weights []weighted, total int) {
	for _, opt := range strings.Split(opts, ";") {
		parts := strings.Split(opt, ":")
		w, err := strconv.ParseInt(parts[len(parts)-1], 10, 64)
		if len(parts) == 1 || err != nil {
			if len(parts) != 1 && err != nil {
				log.Printf("unable to parse option %s defaulting to weight 1: %s", opt, err)
			}
			weights = append(weights, weighted{opt, 1})
			total += 1
			continue
		}
		if w < 0 {
			log.Panicf("weight should be %q > 0: %s", w, opt)
		}
		total += int(w)
		weights = append(weights, weighted{strings.Join(parts[:len(parts)-1], ":"), int(w)})
	}
	if len(weights) == 0 {
		log.Panic("unable to parse even a single option for multiple choice type")
	}
	if total == 0 {
		log.Panic("there should be at least one non-zero weight in the multiple choice options")
	}
	return
}
