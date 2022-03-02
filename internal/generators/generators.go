package generators

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bitstonks/syndi/internal/config"
)

type Generator interface {
	Next() string
}

var generatorBuilders map[string]func(config.ColumnDef) Generator

func RegisterGenerator(genType string, builder func(config.ColumnDef) Generator) {
	generatorBuilders[genType] = builder
}

// GetGenerator will find a generator matching args.Type if one was registered or return an error.
func GetGenerator(args config.ColumnDef) (Generator, error) {
	builder, ok := generatorBuilders[args.Type]
	if !ok {
		return nil, fmt.Errorf("generator of type %s doesn't exist", args.Type)
	}
	return MakeNullifier(builder(args), args.Nullable), nil
}

func init() {
	generatorBuilders = make(map[string]func(config.ColumnDef) Generator)
	// Since the interface always uses strings we can use a unified (weighted) multiple choice random generator
	// for all column types. We do give up on some validation by doing that.
	// TODO: have specialized constructors that validate the data but still use OneOfGenerator behind the scenes
	RegisterGenerator("oneof", NewOneOfGenerator)
	RegisterGenerator("bool/oneof", NewOneOfGenerator)
	RegisterGenerator("datetime/oneof", NewQuotedOneOfGenerator)
	RegisterGenerator("float/oneof", NewOneOfGenerator)
	RegisterGenerator("int/oneof", NewOneOfGenerator)
	RegisterGenerator("string/oneof", NewQuotedOneOfGenerator)

	RegisterGenerator("int/incremental-uniform", NewIntUniformIncrementalGenerator)

	RegisterGenerator("bool", NewBoolGenerator)
	RegisterGenerator("datetime", NewDatetimeNowGenerator)
	RegisterGenerator("datetime/now", NewDatetimeNowGenerator)
	RegisterGenerator("datetime/uniform", NewDatetimeUniformGenerator)
	RegisterGenerator("float", NewFloatUniformGenerator)
	RegisterGenerator("float/uniform", NewFloatUniformGenerator)
	RegisterGenerator("float/normal", NewFloatNormalGenerator)
	RegisterGenerator("float/exp", NewFloatExpGenerator)
	RegisterGenerator("int", NewIntUniformGenerator)
	RegisterGenerator("int/uniform", NewIntUniformGenerator)
	RegisterGenerator("string", NewStringGenerator)
	RegisterGenerator("string/rand", NewStringGenerator)
	RegisterGenerator("string/text", NewTextGenerator)
	RegisterGenerator("string/uuid", NewUuidGenerator)
}

// newRng is a proxy for random object generator, so we can monkey patch it in tests to make them deterministic
var newRng = newRngFunc

func newRngFunc() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
