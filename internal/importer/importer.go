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

func prepareColumnGenerators(columnsConfig map[string]config.Args) ([]string, []generators.Generator) {
	sortedColumns := make([]string, 0, len(columnsConfig))
	for col := range columnsConfig {
		sortedColumns = append(sortedColumns, col)
	}
	sort.Strings(sortedColumns)

	gens := make([]generators.Generator, 0, len(columnsConfig))
	for _, col := range sortedColumns {
		genArgs := columnsConfig[col]
		if genArgs.Type == "" {
			log.Panicf("no data type defined for column `%s`", col)
		}
		g, err := generators.GetGenerator(genArgs)
		if err != nil {
			log.Panic(err)
		}
		gens = append(gens, g)
	}

	return sortedColumns, gens
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
