package generators

import (
	"fmt"
	"github.com/bitstonks/syndi/internal/config"
)

func ExampleNewTextGenerator() {
	args := config.ColumnDef{
		Type:   "string/text",
		Length: 50,
	}
	g, _ := GetGenerator(args)
	fmt.Println(g.Next())
	// Output: 'esque lorem, sit amet malesuada quam consequat qui'
}
