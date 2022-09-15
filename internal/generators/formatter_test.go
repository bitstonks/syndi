package generators

import (
	"github.com/bitstonks/syndi/internal/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFormatter(t *testing.T) {
	g := NewOneOfGenerator(config.ColumnDef{
		Type:  "oneof",
		OneOf: "test",
	})
	t.Run("default value", func(t *testing.T) {
		f := NewFormatter(g, "")
		assert.Equal(t, "test", f.Next())
	})
	t.Run("simple", func(t *testing.T) {
		f := NewFormatter(g, "%v")
		assert.Equal(t, "test", f.Next())
	})
	t.Run("quoted", func(t *testing.T) {
		f := NewFormatter(g, "'%v'")
		assert.Equal(t, "'test'", f.Next())
	})
	t.Run("text", func(t *testing.T) {
		f := NewFormatter(g, "The value is '%v'")
		assert.Equal(t, "The value is 'test'", f.Next())
	})
}

func TestDatetimeDefaultFormatter(t *testing.T) {
	g := NewDatetimeUniformGenerator(config.ColumnDef{
		MaxVal: "2006-01-02 15:04:05",
	})
	f := NewFormatter(g, "")
	assert.Equal(t, "'1971-07-03 10:49:54'", f.Next())
}

func TestDateUSFormatter(t *testing.T) {
	g := NewDatetimeUniformGenerator(config.ColumnDef{
		MaxVal: "2006-01-02 15:04:05",
	})
	f := NewFormatter(g, "Never forget 01/02/06")
	assert.Equal(t, "Never forget 07/03/71", f.Next())
}
