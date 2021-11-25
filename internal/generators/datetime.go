package generators

import (
	"fmt"
	"github.com/bitstonks/syndi/internal/config"
	"log"
	"math/rand"
	"time"
)

func NewDatetimeNowGenerator(args config.Args) Generator {
	args.OneOf = "NOW()"
	return NewOneOfGenerator(args)
}

type DatetimeUniformGenerator struct {
	rng      *rand.Rand
	dtFmt    string
	nullable float64
	minVal   int64
	maxVal   int64
}

func NewDatetimeUniformGenerator(args config.Args) Generator {
	g := DatetimeUniformGenerator{
		rng:      rand.New(rand.NewSource(time.Now().UnixNano())),
		dtFmt:    "2006-01-02 15:04:05",
		nullable: args.Nullable,
	}
	minVal := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC)
	maxVal := time.Now().UTC()
	g.minVal = parseDT(g.dtFmt, args.MinVal, minVal).Unix()
	g.maxVal = parseDT(g.dtFmt, args.MaxVal, maxVal).Unix()
	if g.minVal >= g.maxVal {
		log.Panicf(
			"minVal not smaller than maxVal: %s < %s",
			time.Unix(g.minVal, 0).Format(g.dtFmt),
			time.Unix(g.maxVal, 0).Format(g.dtFmt),
		)
	}
	return &g
}

func (g *DatetimeUniformGenerator) Next() string {
	if g.nullable > 0 && g.rng.Float64() < g.nullable {
		return "NULL"
	}

	secs := g.rng.Int63n(g.maxVal-g.minVal) + g.minVal
	return fmt.Sprintf("%q", time.Unix(secs, 0).Format(g.dtFmt))
}

func parseDT(dtFmt, dt string, fallback time.Time) time.Time {
	if len(dt) == 0 {
		return fallback
	}
	v, err := time.Parse(dtFmt, dt)
	if err != nil {
		log.Panicf("error parsing datetime %q: %s", dt, err)
	}
	return v
}
