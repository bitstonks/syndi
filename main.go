package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"sort"
	"strings"

	"bitbts.bitstamp.net/bint/go-kit/config"
	"bitbts.bitstamp.net/bint/go-kit/version"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	DbDSN        string            `yaml:"DbDSN" validate:"required"`
	DbTable      string            `yaml:"DbTable" validate:"required"`
	TotalRecords int               `yaml:"TotalRecords" validate:"required,gt=0"`
	BatchSize    int               `yaml:"BatchSize" validate:"required,gt=0"`
	SafeImport   bool              `yaml:"SafeImport"`
	Columns      map[string]string `yaml:"Columns" validate:"required,dive,keys,required,endkeys"`
}

// TODO 1: generate in one goroutine and import in another...
// TODO 2: have concurrent clients for faster import...
func main() {
	// parse flags
	configFile := flag.String("c", "config.yaml", "configuration file to use")
	showVersion := flag.Bool("version", false, "show application version and exit immediately")
	flag.Parse()

	if *showVersion {
		fmt.Println(version.GetVersion())
		return
	}

	// load configuration
	var cfg Config
	err := config.LoadAndValidateConfiguration([]string{*configFile}, &cfg)
	if err != nil {
		log.Panicf("error loading config file (%s): %#v:", *configFile, err)
	}

	// connect to db
	db, err := sql.Open("mysql", cfg.DbDSN)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}

	if !cfg.SafeImport {
		// TODO: does this even work?
		log.Println("disabling FK checks")
		_, err = db.Exec("SET FOREIGN_KEY_CHECKS=0")
		if err != nil {
			log.Panic(err)
		}
	}

	cols, gens := prepareColumnGenerators(cfg.Columns)
	sqlPrefix := fmt.Sprintf("INSERT INTO %s (%s) VALUES ", cfg.DbTable, strings.Join(cols, ","))

	var mul, rem int
	// TODO: can i use validation for this?
	if cfg.TotalRecords > cfg.BatchSize {
		mul = cfg.TotalRecords / cfg.BatchSize
		rem = cfg.TotalRecords % cfg.BatchSize
	} else {
		mul = 1
		rem = 0 // for clarity
	}

	for i := 1; i <= mul; i++ {
		log.Printf("loading batch %d of %d records", i, cfg.BatchSize)
		batch := generateBatch(cfg.BatchSize, &gens)
		sql := sqlPrefix + strings.Join(batch, ",")
		_, err = db.Exec(sql)
		if err != nil {
			log.Panic(err)
		}
	}

	if rem != 0 {
		log.Printf("loading remainder of %d records", rem)
		batch := generateBatch(rem, &gens)
		sql := sqlPrefix + strings.Join(batch, ",")
		_, err = db.Exec(sql)
		if err != nil {
			log.Panic(err)
		}
	}

	if !cfg.SafeImport {
		log.Println("enabling FK checks")
		_, err = db.Exec("SET FOREIGN_KEY_CHECKS=1")
		if err != nil {
			log.Panic(err)
		}
	}
}

func prepareColumnGenerators(columnsConfig map[string]string) ([]string, []Generator) {
	sortedColumns := make([]string, 0, len(columnsConfig))
	for col := range columnsConfig {
		sortedColumns = append(sortedColumns, col)
	}
	sort.Strings(sortedColumns)

	generators := make([]Generator, 0, len(columnsConfig))
	for _, col := range sortedColumns {
		genType, genArgs := parseAndValidateDataDefinition(columnsConfig[col])
		if genType == "" {
			log.Panicf("no data type defined for column `%s` (%s)", col, columnsConfig[col])
		}
		switch genType {
		case "bool":
			generators = append(generators, NewBoolGenerator(col, genArgs))
		case "datetime":
			generators = append(generators, NewDatetimeGenerator(col, genArgs))
		case "float":
			generators = append(generators, NewFloatGenerator(col, genArgs))
		case "int":
			generators = append(generators, NewIntGenerator(col, genArgs))
		case "string":
			generators = append(generators, NewStringGenerator(col, genArgs))
		case "text":
			generators = append(generators, NewTextGenerator(col, genArgs))
		case "uuid":
			generators = append(generators, NewUuidGenerator(col, genArgs))
		default:
			log.Panicf("unknown type `%s` in column `%s` (%s)", genType, col, columnsConfig[col])
		}
	}

	return sortedColumns, generators
}

func parseAndValidateDataDefinition(columnDef string) (string, map[string]string) {
	rawDefs := strings.Split(columnDef, ";")

	var genType string
	args := make(map[string]string)

	for _, def := range rawDefs {
		tmp := strings.Split(def, "=")
		if len(tmp) > 2 {
			log.Panicf("invalid column definition: `%s`", def)
		}
		if tmp[0] == "type" {
			genType = tmp[1]
		} else {
			if _, exists := args[tmp[0]]; exists {
				// TODO: consider having the column name here as well
				log.Panicf("multiple definitions for key `%s` in `%s`", tmp[0], def)
			}
			args[tmp[0]] = tmp[1]
		}
	}

	return genType, args
}

func generateBatch(size int, generators *[]Generator) []string {
	res := make([]string, 0, size)
	genSize := len(*generators)

	for i := 0; i < size; i++ {
		single := make([]string, 0, genSize)
		for j := 0; j < genSize; j++ {
			single = append(single, (*generators)[j].Next())
		}
		res = append(res, "("+strings.Join(single, ",")+")")
	}

	return res
}
