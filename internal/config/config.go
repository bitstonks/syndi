package config

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

// RunArgs is a container for command-line flags passed in.
type RunArgs struct {
	Database string `validate:"required"`
	Host     string `validate:"required"`
	Password string `validate:"required"`
	Port     string `validate:"required,number,gt=0"`
	Safe     bool
	Tables   []string `validate:"required,gt=0"`
	User     string   `validate:"required"`
}

func (a RunArgs) GetDSN() string {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&interpolateParams=true",
		a.User,
		a.Password,
		a.Host,
		a.Port,
		a.Database,
	)
	return dsn
}

type Args struct {
	Type     string  `yaml:"Type" validate:"required"`
	Nullable float64 `yaml:"Nullable" validate:"optional"`
	MinVal   string  `yaml:"MinVal" validate:"optional"`
	MaxVal   string  `yaml:"MaxVal" validate:"optional"`
	OneOf    string  `yaml:"OneOf" validate:"optional"`
	Length   int     `yaml:"Length" validate:"optional"`
}

type Config struct {
	DbDSN        string          `yaml:"DbDSN" validate:"required"`
	DbTable      string          `yaml:"DbTable" validate:"required"`
	TotalRecords int             `yaml:"TotalRecords" validate:"required,gt=0"`
	BatchSize    int             `yaml:"BatchSize" validate:"required,gt=0"`
	SafeImport   bool            `yaml:"SafeImport"`
	Columns      map[string]Args `yaml:"Columns" validate:"required,dive,keys,required,endkeys"`
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
