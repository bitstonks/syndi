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

var generatorBuilders map[string]func(config.Args) Generator

func RegisterGenerator(genType string, builder func(config.Args) Generator) {
	generatorBuilders[genType] = builder
}

func GetGenerator(args config.Args) (Generator, error) {
	builder, ok := generatorBuilders[args.Type]
	if !ok {
		return nil, fmt.Errorf("generator of type %s doesn't exist", args.Type)
	}
	return MakeNullable(builder(args), args.Nullable), nil
}

func init() {
	// Since the interface always uses strings we can use a unified (weighted) multiple choice random generator
	// for all column types. We do give up on some validation by doing that.
	// TODO: have specialized constructors that validate the data but still use OneOfGenerator behind the scenes
	RegisterGenerator("oneof", NewOneOfGenerator)
	RegisterGenerator("bool/oneof", NewOneOfGenerator)
	RegisterGenerator("datetime/oneof", NewQuotedOneOfGenerator)
	RegisterGenerator("float/oneof", NewOneOfGenerator)
	RegisterGenerator("int/oneof", NewOneOfGenerator)
	RegisterGenerator("string/oneof", NewQuotedOneOfGenerator)

	RegisterGenerator("bool", NewBoolGenerator)
	RegisterGenerator("datetime", NewDatetimeNowGenerator)
	RegisterGenerator("datetime/now", NewDatetimeNowGenerator)
	RegisterGenerator("datetime/uniform", NewDatetimeUniformGenerator)
	RegisterGenerator("float", NewFloatUniformGenerator)
	RegisterGenerator("float/uniform", NewFloatUniformGenerator)
	RegisterGenerator("float/normal", NewFloatNormalGenerator)
	RegisterGenerator("int", NewIntUniformGenerator)
	RegisterGenerator("int/uniform", NewIntUniformGenerator)
	RegisterGenerator("string", NewStringGenerator)
	RegisterGenerator("string/rand", NewStringGenerator)
	RegisterGenerator("string/text", NewTextGenerator)
	RegisterGenerator("string/uuid", NewUuidGenerator)
}

func NewRng() *rand.Rand {
	return rand.New(rand.NewSource(time.Now().UnixNano()))
}
