[![Go Report Card](https://goreportcard.com/badge/github.com/bitstonks/syndi)](https://goreportcard.com/report/github.com/bitstonks/syndi)

# syndi

syndi is an importer of synthetic data for (My)SQL-type databases. It generates synthetic (artificial and randomized)
data and inserts it into the designated table in the database.

The user describes the type of data they want to have imported for each table column based on the context (and not the 
actual column definition!). For instance, syndi can import incrementing integers that will serve as primary keys. Or it
can choose integers from a range of predefined values that will serve as some code tables.

Data definitions are specified via a YAML config file, one config is (currently) bound to a single database table.

## Available data generators

## Requirements

* Go 1.17+

## TODO

* Unit/integration tests.
* Docker and containerized builds.
* E2E tests.
* Support for other databases.
* Concurrency and/or import speedups.
* Composite columns (data over more than a single table column).
* Composite tables (foreign keys, inheritance/polymorphism).