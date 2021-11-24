package importer

import (
	"database/sql"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/bitstonks/syndi/internal/config"
	"github.com/bitstonks/syndi/internal/generators"
)

type Importer struct {
	db  *sql.DB
	cfg *config.Config
}

func NewImporter(db *sql.DB, cfg *config.Config) *Importer {
	return &Importer{db: db, cfg: cfg}
}

func (im *Importer) DisableFK() error {
	if !im.cfg.SafeImport {
		// TODO: does this even work?
		log.Println("disabling FK checks")
		_, err := im.db.Exec("SET FOREIGN_KEY_CHECKS=0")
		return err
	}
	return nil
}

func (im *Importer) EnableFK() error {
	if !im.cfg.SafeImport {
		log.Println("enabling FK checks")
		_, err := im.db.Exec("SET FOREIGN_KEY_CHECKS=1")
		return err
	}
	return nil
}

func (im *Importer) Import() error {
	cols, gens := prepareColumnGenerators(im.cfg.Columns)
	sqlPrefix := fmt.Sprintf("INSERT INTO %s (%s) VALUES ", im.cfg.DbTable, strings.Join(cols, ","))

	for rem := im.cfg.TotalRecords; rem > 0; rem -= im.cfg.BatchSize {
		log.Printf("loading a batch of max %d out of remaining %d records", im.cfg.BatchSize, rem)
		batch := generateBatch(min(rem, im.cfg.BatchSize), &gens)
		sql := sqlPrefix + strings.Join(batch, ",")
		_, err := im.db.Exec(sql)
		if err != nil {
			return err
		}
	}
	return nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func prepareColumnGenerators(columnsConfig map[string]string) ([]string, []generators.Generator) {
	sortedColumns := make([]string, 0, len(columnsConfig))
	for col := range columnsConfig {
		sortedColumns = append(sortedColumns, col)
	}
	sort.Strings(sortedColumns)

	gens := make([]generators.Generator, 0, len(columnsConfig))
	for _, col := range sortedColumns {
		genType, genArgs := parseAndValidateDataDefinition(columnsConfig[col])
		if genType == "" {
			log.Panicf("no data type defined for column `%s` (%s)", col, columnsConfig[col])
		}
		genArgs["name"] = col
		g, err := generators.GetGenerator(genType, genArgs)
		if err != nil {
			log.Panic(err)
		}
		gens = append(gens, g)
	}

	return sortedColumns, gens
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

func generateBatch(size int, gens *[]generators.Generator) []string {
	res := make([]string, 0, size)
	genSize := len(*gens)

	for i := 0; i < size; i++ {
		single := make([]string, 0, genSize)
		for j := 0; j < genSize; j++ {
			single = append(single, (*gens)[j].Next())
		}
		res = append(res, "("+strings.Join(single, ",")+")")
	}

	return res
}
