package generators

import (
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

//dist=rand;len=20
//dist=oneof;vals=kekec,mojca,rozle
//dist=weights;vals=kekec:20,mojca:10,rozle:2;null=true
type StringGenerator struct {
	rng          *rand.Rand
	Column       string
	Mode         int
	Len          int64
	Nullable     float64
	Values       []string
	Weights      map[string]int64
	WeightsTotal int64
}

func NewStringGenerator(args map[string]string) Generator {
	colName := args["name"]
	g := StringGenerator{
		rng:    rand.New(rand.NewSource(time.Now().UnixNano())),
		Column: colName,
	}
	if v, exists := args["null"]; exists {
		if nullable, err := strconv.ParseFloat(v, 64); err == nil {
			g.Nullable = nullable
		}
	}

	switch args["dist"] {
	case "rand", "":
		if args["dist"] == "" {
			log.Printf("missing `dist` for column %s, defaulting to rand", colName)
		}
		g.Mode = 1
		if lenString, exists := args["len"]; exists {
			l, err := strconv.ParseInt(strings.TrimSpace(lenString), 10, 64)
			if err != nil {
				log.Panicf("error parsing string length %s for column %s ", lenString, colName)
			}
			g.Len = l
		} else {
			log.Printf("absent `len` argument for rand string in column %s, defaulting to 20", colName)
			g.Len = 20
		}

	case "oneof":
		g.Mode = 2
		if _, exists := args["vals"]; !exists {
			log.Panicf("error: `oneof` distribution must include `vals` for column %s", colName)
		}
		g.Values = append(g.Values, strings.Split(args["vals"], ",")...)

	case "weights":
		g.Mode = 3
		if _, exists := args["vals"]; !exists {
			log.Panicf("error: `weights` distribution must include `vals` for column %s", colName)
		}
		// expected: "vals=1:555,2:213,3:102,4:78,5:14,...."
		vals := strings.Split(args["vals"], ",")
		for _, v := range vals {
			valueWeightPair := strings.Split(v, ":")
			weight, err := strconv.ParseInt(strings.TrimSpace(valueWeightPair[1]), 10, 64)
			if err != nil {
				log.Panicf("error parsing weight values from (%s) for column %s", v, colName)
			}
			g.Weights[valueWeightPair[0]] = weight
		}
		for _, v := range g.Weights {
			g.WeightsTotal += v
		}

	default:
		log.Panicf("unknown `dist` %s for column %s", args["dist"], colName)
	}

	return &g
}

// yaay, globals!
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

func (g *StringGenerator) Next() string {
	if g.Nullable > 0 && g.rng.Float64() < g.Nullable {
		return "NULL"
	}
	switch g.Mode {
	case 1:
		return "'" + g.randomString(g.Len) + "'"
	case 2:
		return "'" + g.Values[g.rng.Intn(len(g.Values))] + "'"
	case 3:
		return "'" + g.Values[g.rng.Intn(len(g.Values))] + "'"
	default:
		log.Panicf("unhandled mode `%d` for StringGenerator in column %s", g.Mode, g.Column)
	}
	return ""
}

func (g *StringGenerator) randomString(n int64) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[g.rng.Intn(len(letters))]
	}
	return string(b)
}

// TODO: implement with interfaces{} and remove duplicate code
func (g *StringGenerator) weightedRandomSelect() string {
	r := g.rng.Int63n(g.WeightsTotal)
	for k, v := range g.Weights {
		r -= v
		if r <= 0 {
			return k
		}
	}
	log.Panicf("weightedRandomSelect not really working (r = %d)", r)
	return ""
}
