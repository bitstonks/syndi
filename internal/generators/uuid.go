package generators

import (
	"fmt"
	"github.com/bitstonks/syndi/internal/config"
	"github.com/google/uuid"
	"math/rand"
	"time"
)

type UuidGenerator struct {
	rng      *rand.Rand
	nullable float64
}

// TODO: add length?
func NewUuidGenerator(args config.Args) Generator {
	g := UuidGenerator{
		nullable: args.Nullable,
	}
	if g.nullable > 0 {
		g.rng = rand.New(rand.NewSource(time.Now().UnixNano()))
	}
	return &g
}

func (g *UuidGenerator) Next() string {
	if g.nullable > 0 && g.rng.Float64() < g.nullable {
		return "NULL"
	}
	return fmt.Sprintf("%q", uuid.NewString())
}
