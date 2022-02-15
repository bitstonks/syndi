package generators

import (
	"fmt"
	"log"
	"strconv"
	"sync/atomic"

	"github.com/bitstonks/syndi/internal/config"
)

type incrementalGenerator struct {
	nextVal int64
}

func NewIncrementalGenerator(args config.ColumnDef) Generator {
	minVal, err := strconv.ParseInt(args.MinVal, 10, 64)
	if err != nil {
		log.Panicf("Unable to parse minVal: %s", err)
	}
	return &incrementalGenerator{
		nextVal: minVal,
	}
}

func (g *incrementalGenerator) Next() string {
	return fmt.Sprintf("%d", atomic.AddInt64(&g.nextVal, 1)-1)
}
