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

// ColumnDef defines the type of data we want inserted into a single column of a particular database table.
type ColumnDef struct {
	Type     string  `yaml:"Type" validate:"required"`
	Nullable float64 `yaml:"Nullable" validate:"optional"`
	First    string  `yaml:"First" validate:"optional"`
	MinVal   string  `yaml:"MinVal" validate:"optional"`
	MaxVal   string  `yaml:"MaxVal" validate:"optional"`
	OneOf    string  `yaml:"OneOf" validate:"optional"`
	Length   int     `yaml:"Length" validate:"optional"`
}

// TableDef describes one particular database table. Its data is (mostly) loaded from a YAML file.
type TableDef struct {
	TableName    string               `yaml:"TableName" validate:"required"`
	TotalRecords int                  `yaml:"TotalRecords" validate:"required,gt=0"`
	BatchSize    int                  `yaml:"BatchSize" validate:"required,gt=0"`
	SafeImport   bool                 // TODO: should this be global?
	Columns      map[string]ColumnDef `yaml:"Columns" validate:"required,dive,keys,required,endkeys"`
}

func LoadConfig(args RunArgs) (string, []*TableDef, error) {
	dsn := args.GetDSN()
	tables := make([]*TableDef, 0)
	validate := validator.New()

	err := validate.Struct(args)
	if err != nil {
		reportValidationErrors(err)
		return dsn, tables, err
	}

	// TODO: allow for entire folders to be passed in and implement searching for YAML files in them
	for _, tableFile := range args.Tables {
		yamlFile, err := ioutil.ReadFile(tableFile)
		if err != nil {
			return dsn, tables, err
		}
		tdef := &TableDef{}
		err = yaml.Unmarshal(yamlFile, &tdef)
		if err != nil {
			return dsn, tables, err
		}
		if tdef.BatchSize > tdef.TotalRecords {
			log.Printf("%s: BatchSize larger than TotalRecords, setting the former to equal the latter.\n", tableFile)
			tdef.BatchSize = tdef.TotalRecords
		}
		err = validate.Struct(tdef)
		if err != nil {
			reportValidationErrors(err)
			return dsn, tables, err
		}
		tdef.SafeImport = args.Safe
		tables = append(tables, tdef)
	}

	return dsn, tables, err
}

func reportValidationErrors(err error) {
	if _, ok := err.(*validator.InvalidValidationError); ok {
		fmt.Println(err)
	}
	for _, er := range err.(validator.ValidationErrors) {
		fmt.Println(er)
	}
	fmt.Println()
}
