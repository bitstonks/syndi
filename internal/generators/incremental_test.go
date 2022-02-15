package generators

import (
	"fmt"
	"github.com/bitstonks/syndi/internal/config"
)

func ExampleIncrementalGenerator() {
	args := config.ColumnDef{
		Type:   "int/incremental",
		MinVal: "10", // The first value to be returned
	}
	g, _ := GetGenerator(args)
	fmt.Println(g.Next())
	fmt.Println(g.Next())
	fmt.Println(g.Next())
	// Output:
	// 10
	// 11
	// 12
}
