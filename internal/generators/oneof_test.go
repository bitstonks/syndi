package generators

import (
	"fmt"
	"github.com/bitstonks/syndi/internal/config"
)

func ExampleNewOneOfGenerator_uniform() {
	args := config.Args{
		Type:  "int/oneof",
		OneOf: "1;2;4;8;16",
	}
	g, _ := GetGenerator(args)
	// One of 1, 2, 4, 8, or 16 chosen uniformly at random.
	fmt.Println(g.Next())
	// Output: 16
}

func ExampleNewOneOfGenerator_weighted() {
	args := config.Args{
		Type: "int/oneof",
		// 2 is as likely as 8, but 5 is twice as likely as 2 and 10 three times as likely as 2.
		// Weight '1' is implicitly added to choice 8. And 90 will never be picked.
		OneOf: "2:1;5:2;8;10:3;90:0",
	}
	g, _ := GetGenerator(args)
	// One of 2, 5, 8, or 10 chosen with weighted random.
	fmt.Println(g.Next())
	// Output: 5
}

func ExampleNewQuotedOneOfGenerator_uniform() {
	args := config.Args{
		Type:  "datetime/oneof",
		OneOf: "yes;no",
	}
	g, _ := GetGenerator(args)
	// Either 'yes' or 'no' including single quotes
	fmt.Println(g.Next())
	// Output: 'no'
}

func ExampleNewQuotedOneOfGenerator_weighted() {
	args := config.Args{
		Type:  "string/oneof",
		OneOf: "yes:98;no:1;maybe:1",
	}
	g, _ := GetGenerator(args)
	// Either 'yes' (98%), 'no' (1%), or 'maybe' (1%) including single quotes
	fmt.Println(g.Next())
	// Output: 'yes'
}
