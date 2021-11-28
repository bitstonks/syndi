package generators

import (
	"fmt"

	"github.com/bitstonks/syndi/internal/config"
	"github.com/google/uuid"
)

type UuidGenerator struct{}

// TODO: add length?
func NewUuidGenerator(args config.Args) Generator {
	return &UuidGenerator{}
}

func (g *UuidGenerator) Next() string {
	return fmt.Sprintf("%q", uuid.NewString())
}
