package generators

import (
	"fmt"
	"github.com/bitstonks/syndi/internal/config"
)

func ExampleNewBoolGenerator() {
	args := config.ColumnDef{
		Type: "bool",
	}
	g, _ := GetGenerator(args)
	// One of 1 or 2 chosen uniformly at random.
	fmt.Println(g.Next())
	fmt.Println(g.Next())
	// Output:
	// 1
	// 0
}

func ExampleNewBoolGenerator_oneof() {
	args := config.ColumnDef{
		Type:  "bool/oneof",
		OneOf: "0:20;1:1", // 0 has 20x probability over 1
	}
	g, _ := GetGenerator(args)
	// One of 1 or 2 chosen uniformly at random.
	fmt.Println(g.Next())
	// Output: 0
}
