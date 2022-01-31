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
	db   *sql.DB
	cfg  *config.Config
	cols []string
	gens []generators.Generator
}

func NewImporter(db *sql.DB, cfg *config.Config) *Importer {
	im := Importer{db: db, cfg: cfg}
	im.cols, im.gens = prepareColumnGenerators(cfg.Columns)
	return &im
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
	sqlPrefix := fmt.Sprintf("INSERT INTO %s (%s) VALUES ", im.cfg.DbTable, strings.Join(im.cols, ","))

	for rem := im.cfg.TotalRecords; rem > 0; rem -= im.cfg.BatchSize {
		log.Printf("loading a batch of max %d out of remaining %d records", im.cfg.BatchSize, rem)
		batch := generateBatch(min(rem, im.cfg.BatchSize), &im.gens)
		query := sqlPrefix + strings.Join(batch, ",")
		_, err := im.db.Exec(query)
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

func prepareColumnGenerators(columnsConfig map[string]config.ColumnDef) (cols []string, gens []generators.Generator) {
	for col := range columnsConfig {
		cols = append(cols, col)
	}
	sort.Strings(cols)

	for _, col := range cols {
		genArgs := columnsConfig[col]
		if genArgs.Type == "" {
			log.Panicf("no data type defined for column %q", col)
		}
		g, err := generators.GetGenerator(genArgs)
		if err != nil {
			log.Panic(err)
		}
		gens = append(gens, g)
	}
	return
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
