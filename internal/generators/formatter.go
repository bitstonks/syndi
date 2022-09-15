package generators

import (
	"fmt"
	"time"
)

type DefaultFmt interface {
	DefaultFmtString() string
}

type quotedFmt struct{}

func (quotedFmt) DefaultFmtString() string {
	return "'%v'"
}

type Formatter struct {
	fmtString string
	generator Generator
}

func NewFormatter(g Generator, fmtString string) Generator {
	if fmtString == "" {
		if f, ok := g.(DefaultFmt); ok {
			fmtString = f.DefaultFmtString()
		} else {
			fmtString = "%v"
		}
	}

	return &Formatter{
		fmtString: fmtString,
		generator: g,
	}
}

func (f *Formatter) Next() interface{} {
	val := f.generator.Next()
	if t, ok := val.(time.Time); ok {
		return t.Format(f.fmtString)
	}
	return fmt.Sprintf(f.fmtString, val)
}
