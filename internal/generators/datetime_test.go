package generators

import (
	"fmt"
	"github.com/bitstonks/syndi/internal/config"
)

func ExampleNewDatetimeNowGenerator() {
	args := config.Args{
		Type: "datetime", // alias for "datetime/now"
	}
	g, _ := GetGenerator(args)
	fmt.Println(g.Next())
	// Output: NOW()
}

func ExampleNewDatetimeUniformGenerator() {
	args := config.Args{
		Type:   "datetime/uniform",
		MinVal: "2011-08-15 18:18:18", // Default is 1970-01-01 00:00:00
		MaxVal: "2021-12-01 21:54:35", // Default is current time
	}
	g, _ := GetGenerator(args)
	fmt.Println(g.Next())
	// Output: '2016-12-20 00:42:51'
}
