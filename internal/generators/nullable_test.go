package generators

import (
	"fmt"
	"github.com/bitstonks/syndi/internal/config"
)

func ExampleMakeNullable_half() {
	args := config.Args{
		Type:     "oneof",
		OneOf:    "ok", // only possible non-null value is 'ok'
		Nullable: 0.2,  // NULL 20% of time
	}
	g, _ := GetGenerator(args)
	fmt.Println(g.Next())
	fmt.Println(g.Next())
	// Output:
	// ok
	// NULL
}

func ExampleMakeNullable_always() {
	args := config.Args{
		Type:     "int/uniform", // uniform random number
		MinVal:   "0",           // from including 0
		MaxVal:   "100",         // to not including 100
		Nullable: 1,             // but it will be NULL 100% of time
	}
	g, _ := GetGenerator(args)
	fmt.Println(g.Next())
	// Output: NULL
}
