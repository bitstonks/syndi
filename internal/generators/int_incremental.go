package generators

import (
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

func (g *intUniformIncrementalGenerator) Next() interface{} {
	step := parseInt(g.generator.Next())
	result := atomic.AddInt64(&g.nextValue, step) - step
	return result
}

func parseInt(raw interface{}) int64 {
	switch n := raw.(type) {
	case int:
		return int64(n)
	case int64:
		return n
	case string:
		parsed, err := strconv.ParseInt(n, 10, 64)
		if err != nil {
			log.Panicf("int generator returned invalid int: %s", err)
		}
		return parsed
	default:
		log.Panicf("unable to parse int from %v", n)
		return 0
	}
}
