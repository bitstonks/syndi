package generators

import (
	"fmt"

	"github.com/bitstonks/syndi/internal/config"
	"github.com/google/uuid"
)

// It's here so that it can be monkey patched in tests
var uuidGen = uuid.NewString

type uuidGenerator struct{}

// TODO: add length?
func NewUuidGenerator(args config.ColumnDef) Generator {
	return &uuidGenerator{}
}

func (g *uuidGenerator) Next() string {
	return fmt.Sprintf("'%s'", uuidGen())
}
