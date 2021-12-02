[![Go Report Card](https://goreportcard.com/badge/github.com/bitstonks/syndi)](https://goreportcard.com/report/github.com/bitstonks/syndi)

# syndi

syndi is an importer of synthetic data for (My)SQL-type databases. It generates synthetic (artificial and randomized)
data and inserts it into the designated table in the database.

The user describes the type of data they want to have imported for each table column based on the context (and not the 
actual column definition!). For instance, syndi can import incrementing integers that will serve as primary keys. Or it
can choose integers from a range of predefined values that will serve as some code tables.

Data definitions are specified via a YAML config file, one config is (currently) bound to a single database table.

## Available data generators

> **_NOTE:_** this list was maybe accurate at the time of writing but might not be at the time of you reading it.
> Consider these values as approximate and see
> [examples in go docs](https://pkg.go.dev/github.com/bitstonks/syndi/internal/generators#Generator)
> for more up-to-date information.

|       Type/Parameter        | Description |
| ---------------   | ----------- |
| **Type: bool** | Generates 50% `0` and 50% `1`  |
| ======== | ================= |
| **Type: bool/oneof** | Generates `0` or `1` depending on OneOf weights. |
| OneOf: str | Semicolon separated values with optional weights separated by colons. (e.g. `"1"`, `"0:5;1:95"`) |
| ======== | ================= |
| **Type: datetime, datetime/now** | Generates `NOW()`  |
| ======== | ================= |
| **Type: datetime/uniform** | Dates between MinVal and MaxVal (format: `2006-01-02 15:04:05`). |
| MinVal: date  | Minimal date to generate (default is unix epoch). |
| MaxVal: date  | Maximal date to generate (default is current time). |
| ======== | ================= |
| **Type: datetime/oneof** | One of the given dates (format: `2006-01-02 15:04:05`). |
| OneOf: str | Semicolon separated values with weights separated by colons. <br>**NOTE:** weights are mandatory else seconds will be interpreted as weights resulting in an invalid datetime.|
| ======== | ================= |
| **Type: float, float/uniform** | Generates random floats in [MinVal, MaxVal) |
| MinVal: float  | Minimal value to generate. |
| MaxVal: float  | Maximal value to generate.|
| ======== | ================= |
| **Type: float, float/normal** | Generates floats from \[-math.MaxFloat64,  +math.MaxFloat64] with normal distribution where `mean=(MinVal+MaxVal)/2` and `stDev=(MaxVal-MinVal)/2`. |
| MinVal: float  | Value one standard deviation below the mean. |
| MaxVal: float  | Value one standard deviation above the mean.|
| ======== | ================= |
| **Type: float, float/exp** | Generates floats from (MinVal,  math.MaxFloat64] with exponential distribution with `mean=(MinVal+MaxVal)/2` and ~15% of values are bigger than MaxVal. |
| MinVal: float  | Minimal value to generate. |
| MaxVal: float  | Defines how spread out you want the numbers to be. Mean value will be the average of MinVal and MaxVal.|
| ======== | ================= |
| **Type: float/oneof** | Choose one of the options given. |
| OneOf: str | Semicolon separated values with optional weights separated by colons. (e.g. `"1.5;2.1;4.0"`, `"0.1:3;2.5:10"`) |
| ======== | ================= |
| **Type: int, int/uniform** | Generates random integers in [MinVal, MaxVal) |
| MinVal: int  | Minimal value to generate (inclusive). |
| MaxVal: int  | Maximal value to generate (non-inclusive).|
| ======== | ================= |
| **Type: int/oneof** | Chooses one of the options given. |
| OneOf: str | Semicolon separated values with optional weights separated by colons. (e.g. `"1;2;4;8"`, `"1:3;2:10"`) |
| ======== | ================= |
| **Type: string, string/rand** | Generates a random string of given length. |
| Length: int | Number of characters in the output string. |
| OneOf: str | (Optional) Character set to pick the characters from.<br>Default: `abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789` |
| ======== | ================= |
| **Type: string/text** | Generates random sections of lorem ipsum text of given length. |
| Length: int | Number of characters in the output string. |
| ======== | ================= |
| **Type: string/uuid** | Generates a 36 characters long universally unique string. |
| ======== | ================= |
| **Type: string/oneof** | Chooses one of the options given. |
| OneOf: str | Semicolon separated values with optional weights separated by colons. (e.g. `yes;no`, `yes:10;no:1;maybe:4`) |
| ======== | ================= |
| **Nullable (any type)** | Any type can be made to return part of the results to be `NULL` by setting the Nullable parameter. |
| Nullable: float | A number between 0 and 1 defining the probability of each value to be `NULL`. |

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