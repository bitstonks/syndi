package config

import (
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadConfig(t *testing.T) {
	currWd, err := os.Getwd()
	assert.NoError(t, err)

	cfgPath := path.Join(currWd, "../../test/testdata/config-example.yaml")
	cfg, err := LoadConfig(cfgPath)
	assert.NoError(t, err)
	assert.False(t, cfg.SafeImport)
	assert.Equal(t, 5031, cfg.TotalRecords)
}
