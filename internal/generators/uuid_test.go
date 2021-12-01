package generators

import (
	"fmt"
	"github.com/bitstonks/syndi/internal/config"
)

func init() {
	uuidGen = func() string {
		return "4d618232-ae05-46d0-a270-2931ef3d9add"
	}
}

func ExampleNewUuidGenerator() {
	args := config.Args{
		Type: "string/uuid",
	}
	g, _ := GetGenerator(args)
	fmt.Println(g.Next())
	// Output: '4d618232-ae05-46d0-a270-2931ef3d9add'
}
