package generators

import (
	"fmt"
	"github.com/google/uuid"
	"math/rand"
	"strconv"
	"time"
)

type UuidGenerator struct {
	rng      *rand.Rand
	Nullable float64
}

// TODO: add length?
func NewUuidGenerator(args map[string]string) Generator {
	g := UuidGenerator{}
	if v, exists := args["null"]; exists {
		if nullable, err := strconv.ParseFloat(v, 64); err == nil {
			g.Nullable = nullable
		}
		g.rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	return &g
}

func (g *UuidGenerator) Next() string {
	if g.Nullable > 0 && g.rng.Float64() < g.Nullable {
		return "NULL"
	}
	return fmt.Sprintf("%q", uuid.NewString())
}
