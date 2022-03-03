package generators

import (
	"github.com/bitstonks/syndi/internal/config"
	"github.com/google/uuid"
)

// It's here so that it can be monkey patched in tests
var uuidGen = uuid.NewString

type uuidGenerator struct {
	quotedFmt
}

// TODO: add length?
func NewUuidGenerator(args config.ColumnDef) Generator {
	return &uuidGenerator{}
}

func (g *uuidGenerator) Next() interface{} {
	return uuidGen()
}
