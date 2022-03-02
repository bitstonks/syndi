package generators

import (
	"fmt"
	"github.com/bitstonks/syndi/internal/config"
)

func ExampleNewIntUniformIncrementalGenerator() {
	args := config.ColumnDef{
		Type:   "int/incremental-uniform",
		First:  "0",  // start with this
		MinVal: "1",  // increase by at least
		MaxVal: "20", // increase less than
	}
	g, _ := GetGenerator(args)
	// Starting at 0 it increases numbers by [1, 20) every step.
	fmt.Println(g.Next())
	fmt.Println(g.Next())
	fmt.Println(g.Next())
	// Output:
	// 0
	// 11
	// 26
}
