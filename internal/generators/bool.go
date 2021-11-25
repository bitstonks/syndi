package generators

import (
	"github.com/bitstonks/syndi/internal/config"
)

func NewBoolGenerator(args config.Args) Generator {
	args.OneOf = "0;1"
	return NewOneOfGenerator(args)
}
