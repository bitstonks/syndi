package generators

import (
	"fmt"
	"github.com/bitstonks/syndi/internal/config"
)

func ExampleNewStringGenerator() {
	args := config.ColumnDef{
		Type:   "string", // alias for string/rand
		Length: 15,
	}
	g, _ := GetGenerator(args)
	fmt.Println(g.Next())
	// Output: 'rgltBHYVJQVAdv8'
}

func ExampleNewStringGenerator_charset() {
	args := config.ColumnDef{
		Type:   "string/rand",
		Length: 5,
		OneOf:  " abcd", // choose the charset you want to pick from
	}
	g, _ := GetGenerator(args)
	fmt.Println(g.Next())
	// Output: 'dac b'
}
