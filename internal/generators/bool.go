package generators

import (
	"math/rand"
	"strconv"
	"time"
)

type BoolGenerator struct {
	rng      *rand.Rand
	Column   string
	Nullable float64
}

func NewBoolGenerator(args map[string]string) Generator {
	colName := args["name"]
	g := BoolGenerator{
		rng:    rand.New(rand.NewSource(time.Now().UnixNano())),
		Column: colName,
	}
	if v, exists := args["null"]; exists {
		if nullable, err := strconv.ParseFloat(v, 64); err == nil {
			g.Nullable = nullable
		}
	}
	return &g
}

func (g *BoolGenerator) Next() string {
	if g.Nullable > 0 && g.rng.Float64() < g.Nullable {
		return "NULL"
	}
	if g.rng.Float32() < 0.5 {
		return "0"
	} else {
		return "1"
	}
}
