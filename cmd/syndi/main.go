package main

import (
	"database/sql"
	"flag"
	"log"

	"github.com/bitstonks/syndi/internal/config"
	"github.com/bitstonks/syndi/internal/importer"

	_ "github.com/go-sql-driver/mysql"
)

// TODO 1: generate in one goroutine and import in another...
// TODO 2: have concurrent clients for faster import...
func main() {
	// save command-line arguments
	args := config.RunArgs{}
	flag.StringVar(&args.Database, "db", "bitstamp_dev", "Database name to use")
	flag.StringVar(&args.Host, "host", "localhost", "Database host to connect to")
	flag.StringVar(&args.Password, "p", "root", "Database user's password")
	flag.StringVar(&args.Port, "P", "28000", "Database port number")
	flag.BoolVar(&args.Safe, "safe", false, "Whether foreign key checks are mandated")
	flag.StringVar(&args.User, "u", "root", "Database user")
	flag.Parse()
	args.Tables = flag.Args()

	// load configuration
	dbDSN, tableDefinitions, err := config.LoadConfig(args)
	if err != nil {
		log.Panicf("error loading config: %#v:", err)
	}

	// connect to db
	db, err := sql.Open("mysql", dbDSN)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Panic(err)
	}

	// import things
	for _, tableDef := range tableDefinitions {
		im := importer.NewImporter(db, tableDef)
		err = im.DisableFK()
		if err != nil {
			log.Panic(err)
		}

		err = im.Import()
		if err != nil {
			log.Panic(err)
		}
		im.EnableFK() // TODO: this can now fail to run, not sure whether it is a problem since it's connection-bound (?)
	}
}
