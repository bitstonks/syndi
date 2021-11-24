package generators

import (
	"fmt"
)

type Generator interface {
	Next() string
}

var generatorBuilders map[string]func(map[string]string) Generator

func RegisterGenerator(genType string, builder func(map[string]string) Generator) {
	generatorBuilders[genType] = builder
}

func GetGenerator(genType string, args map[string]string) (Generator, error) {
	builder, ok := generatorBuilders[genType]
	if !ok {
		return nil, fmt.Errorf("generator of type %s doesn't exist", genType)
	}
	return builder(args), nil
}

func init() {
	RegisterGenerator("bool", NewBoolGenerator)
	RegisterGenerator("datetime", NewDatetimeGenerator)
	RegisterGenerator("float", NewFloatGenerator)
	RegisterGenerator("int", NewIntGenerator)
	RegisterGenerator("string", NewStringGenerator)
	RegisterGenerator("text", NewTextGenerator)
	RegisterGenerator("uuid", NewUuidGenerator)
}
