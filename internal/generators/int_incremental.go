package generators

import (
	"fmt"
	"github.com/bitstonks/syndi/internal/config"
	"log"
	"strconv"
	"sync/atomic"
)

type intUniformIncrementalGenerator struct {
	nextValue int64
	generator Generator
}

func NewIntUniformIncrementalGenerator(args config.ColumnDef) Generator {
	first, err := strconv.ParseInt(args.First, 10, 64)
	if err != nil {
		log.Panicf("Unable to parse first: %s", err)
	}
	return &intUniformIncrementalGenerator{
		nextValue: first,
		generator: NewIntUniformGenerator(args),
	}
}

func (g *intUniformIncrementalGenerator) Next() string {
	step, err := strconv.ParseInt(g.generator.Next(), 10, 64)
	if err != nil {
		log.Panicf("int generator returned invalid int: %s", err)
	}
	result := atomic.AddInt64(&g.nextValue, step) - step
	return fmt.Sprintf("%d", result)
}
