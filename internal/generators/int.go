package generators

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type IntGenerator struct {
	rng          *rand.Rand
	Column       string
	Mode         int
	Nullable     float64
	Values       []int64
	Weights      map[int64]int64
	WeightsTotal int64
}

//vals=1,2000;dist=uniform
//vals=1,2000;dist=oneof
//vals=1:2000,2:400,3:3200;dist=weights;null=true
func NewIntGenerator(args map[string]string) Generator {
	colName := args["name"]
	g := IntGenerator{
		rng:     rand.New(rand.NewSource(time.Now().UnixNano())),
		Column:  colName,
		Weights: make(map[int64]int64),
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
			min, err := strconv.ParseInt(strings.TrimSpace(boundaries[0]), 10, 64)
			if err != nil {
				log.Panicf("error parsing min value for column %s in %s", colName, val)
			}
			g.Values = append(g.Values, min)

			max, err := strconv.ParseInt(strings.TrimSpace(boundaries[1]), 10, 64)
			if err != nil {
				log.Panicf("error parsing max value for column %s in %s", colName, val)
			}
			g.Values = append(g.Values, max+1)
		} else {
			log.Printf("absent `vals` for uniform int in column %s, defaulting to [0,100]", colName)
			g.Values = append(g.Values, []int64{0, 101}...)
		}

	case "oneof":
		g.Mode = 2
		if _, exists := args["vals"]; !exists {
			log.Panicf("error: `oneof` distribution must include `vals` for column %s", colName)
		}
		// expected: "vals=n1,n2,n3,..."
		vals := strings.Split(args["vals"], ",")
		for _, v := range vals {
			num, err := strconv.ParseInt(strings.TrimSpace(v), 10, 64)
			if err != nil {
				log.Panicf("error parsing ints from (%s) for column %s", vals, colName)
			}
			g.Values = append(g.Values, num)
		}

	case "weights":
		g.Mode = 3
		if _, exists := args["vals"]; !exists {
			log.Panicf("error: `weights` distribution must include `vals` for column %s", colName)
		}
		// expected: "vals=1:555,2:213,3:102,4:78,5:14,...."
		vals := strings.Split(args["vals"], ",")
		for _, v := range vals {
			valueWeightPair := strings.Split(v, ":")
			num, err := strconv.ParseInt(strings.TrimSpace(valueWeightPair[0]), 10, 64)
			if err != nil {
				log.Panicf("error parsing number values from (%s) for column %s", v, colName)
			}
			weight, err := strconv.ParseInt(strings.TrimSpace(valueWeightPair[1]), 10, 64)
			if err != nil {
				log.Panicf("error parsing weight values from (%s) for column %s", v, colName)
			}
			g.Weights[num] = weight
		}
		for _, v := range g.Weights {
			g.WeightsTotal += v
		}

	default:
		log.Panicf("unknown `dist` %s for column %s", args["dist"], colName)
	}

	return &g
}

func (g *IntGenerator) Next() string {
	if g.Nullable > 0 && g.rng.Float64() < g.Nullable {
		return "NULL"
	}
	var v int64
	switch g.Mode {
	case 1:
		v = g.Values[0] + g.rng.Int63n(g.Values[1]-g.Values[0])
	case 2:
		v = g.Values[g.rng.Intn(len(g.Values)-1)]
	case 3:
		v = g.weightedRandomSelect()
	default:
		log.Panicf("unhandled mode `%d` for IntGenerator in column %s", g.Mode, g.Column)
	}
	return fmt.Sprintf("%d", v)
}

// Stolen here: https://medium.com/@peterkellyonline/weighted-random-selection-3ff222917eb6
// Could be stateless but since the worker already holds state, the API is smaller this way.
func (g *IntGenerator) weightedRandomSelect() int64 {
	r := g.rng.Int63n(g.WeightsTotal)
	for k, v := range g.Weights {
		r -= v
		if r <= 0 {
			return k
		}
	}
	log.Panicf("weightedRandomSelect not really working (r = %d)", r)
	return 0
}
