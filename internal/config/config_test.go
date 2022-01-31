package config

import (
	"github.com/go-playground/validator/v10"
	"os"
	"path"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunArgs(t *testing.T) {
	t.Run("test initialization happy path", func(t *testing.T) {
		args := RunArgs{
			Database: "example",
			Host:     "localhost",
			Password: "root",
			Port:     "3306",
			Safe:     false,
			Tables:   []string{"users.yaml", "accounts.yaml"},
			User:     "root",
		}
		validate := validator.New()
		err := validate.Struct(args)
		assert.NoError(t, err)
	})

	t.Run("test initialization sad path", func(t *testing.T) {
		args := RunArgs{
			Database: "example",
			Host:     "localhost",
			Password: "root",
			Port:     "invalid port number",
			Safe:     false,
			User:     "root",
		}
		validate := validator.New()
		err := validate.Struct(args)
		assert.EqualError(t, err, "Key: 'RunArgs.Port' Error:Field validation for 'Port' failed on the 'number' tag\nKey: 'RunArgs.Tables' Error:Field validation for 'Tables' failed on the 'required' tag")

		args.Port = "3306"
		err = validate.Struct(args)
		assert.EqualError(t, err, "Key: 'RunArgs.Tables' Error:Field validation for 'Tables' failed on the 'required' tag")

		args.Tables = []string{"users.yaml"}
		err = validate.Struct(args)
		assert.NoError(t, err)
	})

	t.Run("test GetDSN", func(t *testing.T) {
		args := RunArgs{
			Database: "example",
			Host:     "localhost",
			Password: "root",
			Port:     "3306",
			Safe:     false,
			Tables:   []string{"users.yaml", "accounts.yaml"},
			User:     "root",
		}
		assert.Equal(t, args.GetDSN(), "root:root@tcp(localhost:3306)/example?parseTime=true&interpolateParams=true")
	})
}

func TestLoadConfig(t *testing.T) {
	currWd, err := os.Getwd()
	assert.NoError(t, err)

	cfgPath := path.Join(currWd, "../../test/testdata/config-example.yaml")
	cfg, err := LoadConfig(cfgPath)
	assert.NoError(t, err)
	assert.False(t, cfg.SafeImport)
	assert.Equal(t, 5031, cfg.TotalRecords)
}
