package generators

import (
	"log"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type DatetimeGenerator struct {
	rng      *rand.Rand
	Column   string
	DttmFmt  string
	Mode     int
	Nullable float64
	Values   []int64
}

func NewDatetimeGenerator(args map[string]string) Generator {
	colName := args["name"]
	g := DatetimeGenerator{
		rng:     rand.New(rand.NewSource(time.Now().UnixNano())),
		Column:  colName,
		DttmFmt: "2006-01-02 15:04:05",
	}
	if v, exists := args["null"]; exists {
		if nullable, err := strconv.ParseFloat(v, 64); err == nil {
			g.Nullable = nullable
		}
	}

	switch args["dist"] {
	case "now", "":
		if args["dist"] == "" {
			log.Printf("missing `dist` for column %s, defaulting to `now`", colName)
		}
		g.Mode = 1

	case "rand":
		g.Mode = 2
		if strBounds, exists := args["vals"]; exists {
			bounds := strings.Split(strBounds, ",")
			t, err := time.Parse(g.DttmFmt, bounds[0])
			if err != nil {
				log.Panicf("error parsing datetime bound `%s` for column %s", bounds[0], colName)
			}
			g.Values = append(g.Values, t.Unix())

			// TODO: mulitple values?!
			if len(bounds) == 2 {
				t, err = time.Parse(g.DttmFmt, bounds[1])
				if err != nil {
					log.Panicf("error parsing datetime bound `%s` for column %s", bounds[1], colName)
				}
				g.Values = append(g.Values, t.Unix())
			} else {
				g.Values = append(g.Values, time.Now().UTC().Unix())
			}
		} else {
			g.Values = []int64{time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix(), time.Now().UTC().Unix()}
		}

	default:
		log.Panicf("unknown `dist` %s for column %s", args["dist"], colName)
	}
	return &g
}

func (g *DatetimeGenerator) Next() string {
	if g.Nullable > 0 && g.rng.Float64() < g.Nullable {
		return "NULL"
	}

	var res string
	switch g.Mode {
	case 1:
		res = "NOW()"
	case 2:
		secs := g.rng.Int63n(g.Values[1]-g.Values[0]) + g.Values[0]
		res = "'" + time.Unix(secs, 0).Format(g.DttmFmt) + "'"
	}
	return res
}
