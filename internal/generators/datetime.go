package generators

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/bitstonks/syndi/internal/config"
)

func NewDatetimeNowGenerator(args config.Args) Generator {
	args.OneOf = "NOW()"
	return NewOneOfGenerator(args)
}

type datetimeUniformGenerator struct {
	rng    *rand.Rand
	dtFmt  string
	minVal int64
	spread int64
}

func NewDatetimeUniformGenerator(args config.Args) Generator {
	g := datetimeUniformGenerator{
		rng:   newRng(),
		dtFmt: "2006-01-02 15:04:05",
	}
	minVal := time.Date(1970, 1, 0, 0, 0, 0, 0, time.UTC).Unix()
	maxVal := time.Now().UTC().Unix()
	minVal = parseDT(g.dtFmt, args.MinVal, minVal)
	maxVal = parseDT(g.dtFmt, args.MaxVal, maxVal)
	if minVal >= maxVal {
		log.Panicf(
			"minVal not smaller than maxVal: %s < %s",
			time.Unix(minVal, 0).Format(g.dtFmt),
			time.Unix(maxVal, 0).Format(g.dtFmt),
		)
	}
	g.minVal = minVal
	g.spread = maxVal - minVal
	return &g
}

func (g *datetimeUniformGenerator) Next() string {
	secs := g.rng.Int63n(g.spread) + g.minVal
	return fmt.Sprintf("'%s'", time.Unix(secs, 0).UTC().Format(g.dtFmt))
}

func parseDT(dtFmt, dt string, fallback int64) int64 {
	if len(dt) == 0 {
		return fallback
	}
	v, err := time.Parse(dtFmt, dt)
	if err != nil {
		log.Panicf("error parsing datetime %q: %s", dt, err)
	}
	return v.Unix()
}
