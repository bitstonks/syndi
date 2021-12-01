package generators

import (
	"fmt"
	"github.com/bitstonks/syndi/internal/config"
)

func ExampleNewIntUniformGenerator() {
	args := config.Args{
		Type:   "int", // alias for "int/uniform"
		MinVal: "-50", // inclusive
		MaxVal: "202", // non-inclusive
	}
	g, _ := GetGenerator(args)
	// Returns a number in [-50, 202) uniformly at random
	fmt.Println(g.Next())
	// Output: 119
}
