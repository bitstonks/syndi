package config

import "io/ioutil"
import "gopkg.in/yaml.v2"

// TODO(mkolman): Validate using github.com/go-playground/validator/v10

type Config struct {
	DbDSN        string            `yaml:"DbDSN" validate:"required"`
	DbTable      string            `yaml:"DbTable" validate:"required"`
	TotalRecords int               `yaml:"TotalRecords" validate:"required,gt=0"`
	BatchSize    int               `yaml:"BatchSize" validate:"required,gt=0"`
	SafeImport   bool              `yaml:"SafeImport"`
	Columns      map[string]string `yaml:"Columns" validate:"required,dive,keys,required,endkeys"`
}

func LoadConfig(filename *string) (*Config, error) {
	yamlFile, err := ioutil.ReadFile(*filename)
	if err != nil {
		return nil, err
	}
	c := &Config{}
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		return nil, err
	}
	return c, nil
}
