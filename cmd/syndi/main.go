package main

import (
	"database/sql"
	"flag"
	"log"

	"github.com/bitstonks/syndi/internal/config"
	"github.com/bitstonks/syndi/internal/importer"

	_ "github.com/go-sql-driver/mysql"
)

var configFile = flag.String("c", "config.yaml", "configuration file to use")

// TODO 1: generate in one goroutine and import in another...
// TODO 2: have concurrent clients for faster import...
func main() {
	flag.Parse()

	// load configuration
	cfg, err := config.LoadConfig(*configFile)
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

	im := importer.NewImporter(db, cfg)
	err = im.DisableFK()
	if err != nil {
		log.Panic(err)
	}
	defer im.EnableFK()

	err = im.Import()
	if err != nil {
		log.Panic(err)
	}
}
