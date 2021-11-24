package generators

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type FloatGenerator struct {
	rng      *rand.Rand
	Column   string
	Mode     int
	Nullable float64
	Values   []float64
}

// vals=1,2000;dist=uniform
func NewFloatGenerator(args map[string]string) Generator {
	colName := args["name"]
	g := FloatGenerator{
		rng:    rand.New(rand.NewSource(time.Now().UnixNano())),
		Column: colName,
	}
	if v, exists := args["null"]; exists {
		if nullable, err := strconv.ParseFloat(v, 64); err == nil {
			g.Nullable = nullable
		}
	}
	switch args["dist"] {
	case "uniform", "":
		if args["dist"] == "" {
			log.Printf("missing `dist` for column %s, defaulting to uniform", colName)
		}
		g.Mode = 1
		// expected: "vals=min,max"
		if val, exists := args["vals"]; exists {
			boundaries := strings.Split(val, ",")
			// TODO: solve len(boundaries) != 2 !!!
			min, err := strconv.ParseFloat(strings.TrimSpace(boundaries[0]), 64)
			if err != nil {
				log.Panicf("error parsing min value for column %s in %s", colName, val)
			}
			g.Values = append(g.Values, min)

			max, err := strconv.ParseFloat(strings.TrimSpace(boundaries[1]), 64)
			if err != nil {
				log.Panicf("error parsing max value for column %s in %s", colName, val)
			}
			g.Values = append(g.Values, max)
		} else {
			log.Printf("absent `vals` for uniform int in column %s, defaulting to [0,100]", colName)
			g.Values = append(g.Values, []float64{0.0, 100.0}...)
		}

	case "normal":
		g.Mode = 2

	default:
		log.Panicf("unknown `dist` %s for column %s", args["dist"], colName)
	}

	return &g
}

func (g *FloatGenerator) Next() string {
	if g.Nullable > 0 && g.rng.Float64() < g.Nullable {
		return "NULL"
	}
	var v float64
	switch g.Mode {
	case 1:
		v = g.Values[0] + g.rng.Float64()*(g.Values[1]-g.Values[0])
	case 2:
		v = g.rng.NormFloat64() // TODO: better normal dist!
	default:
		log.Panicf("unhandled mode `%d` for FloatGenerator in column %s", g.Mode, g.Column)
	}
	return fmt.Sprintf("%g", v) // TODO: is the format correct at all?
}
