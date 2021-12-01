package generators

import (
	"fmt"
	"github.com/bitstonks/syndi/internal/config"
)

func ExampleNewFloatUniformGenerator() {
	args := config.Args{
		Type:   "float", // alias for "float/uniform"
		MinVal: "-1.5",
		MaxVal: "0.36",
	}
	g, _ := GetGenerator(args)
	// Returns a number in [-1.5, 0.36) uniformly at random
	fmt.Println(g.Next())
	// Output: -1.0473285656995142
}

func ExampleNewFloatNormalGenerator() {
	args := config.Args{
		Type:   "float/normal",
		MinVal: "10.0", // mean - stdev
		MaxVal: "30.0", // mean + stdev
	}
	g, _ := GetGenerator(args)
	// Returns any number distributed normally with mean=20 and stddev=10
	fmt.Println(g.Next())
	// Output: 27.133352143941206
}

func ExampleNewFloatExpGenerator() {
	args := config.Args{
		Type:   "float/exp",
		MinVal: "10.0",
		MaxVal: "30.0", // MinVal + 1/lambda
	}
	g, _ := GetGenerator(args)
	// Returns a number greater than 10 falling exponentially with mean=20
	// Around 15% of numbers will be bigger than 30.
	fmt.Println(g.Next())
	// Output: 12.181511239890659
}
