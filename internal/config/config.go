package config

import (
	"io/ioutil"
	"log"

	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"
)

type Config struct {
	DbDSN        string            `yaml:"DbDSN" validate:"required"`
	DbTable      string            `yaml:"DbTable" validate:"required"`
	TotalRecords int               `yaml:"TotalRecords" validate:"required,gt=0"`
	BatchSize    int               `yaml:"BatchSize" validate:"required,gt=0"`
	SafeImport   bool              `yaml:"SafeImport"`
	Columns      map[string]string `yaml:"Columns" validate:"required,dive,keys,required,endkeys"`
}

func LoadConfig(filename string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return nil, err
	}
	if c.BatchSize > c.TotalRecords {
		log.Println("Setting BatchSize to equal TotalRecords.")
		c.BatchSize = c.TotalRecords
	}
	validate := validator.New()
	err = validate.Struct(c)
	return c, err
}
