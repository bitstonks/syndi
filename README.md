[![Go Report Card](https://goreportcard.com/badge/github.com/bitstonks/syndi)](https://goreportcard.com/report/github.com/bitstonks/syndi)

# syndi

syndi is an importer of synthetic data for (My)SQL-type databases. It generates synthetic (artificial and randomized)
data and inserts it into the designated table in the database.

The user describes the type of data they want to have imported for each table column based on the context (and not the 
actual column definition!). For instance, syndi can import incrementing integers that will serve as primary keys. Or it
can choose integers from a range of predefined values that will serve as some code tables.

Data definitions are specified via a YAML config file, one config is (currently) bound to a single database table.

## Usage

Clone the repo and compile the binary. The `-h` command-line flag hints at possible parameters
```shell
$ ./syndi -h
...
```

You must pass at least one table definition YAML file, for example.
```shell
$ ./syndi users.yaml
```

See `test/testdata/config-example.yaml` and generator definitions in the next paragraph for details.

## Available data generators

> **_NOTE:_** this list was up-to-date at the time of writing but might not be at the time of you reading it.
> Consider these values as approximation and see
> [examples in go docs](https://pkg.go.dev/github.com/bitstonks/syndi/internal/generators#pkg-examples)
> for more accurate information.

Below are examples of configuration for different kind of generators. These can be used as values in the `Columns`
section of your config. For full config structure you can check out the
[docs on config.Config](https://pkg.go.dev/github.com/bitstonks/syndi/internal/config#Config).

```yaml
bool1:
  # Generates 50% `0` and 50% `1`.
  Type: bool
bool2:
  # Generates `0` or `1` depending on OneOf weights.
  Type: bool/oneof
  OneOf: 0:1;1:30  # `1` is 30 times more probable than `0`.
datetime1:
  # Generates `NOW()`.
  Type: datetime  # Alias for `datetime/now`.
datetime2:
  # Generates random dates uniformly at random from [MinVal, MaxVal).
  Type: datetime/uniform
  MinVal: 2011-08-15 18:18:18  # Default is `1970-01-01 00:00:00`.
  MaxVal: 2021-12-01 21:54:35  # Default is current time.
datetime3:
  # Selects one of the dates given in OneOf.
  Type: datetime/oneof
  # Weights are mandatory, because dates include colons.
  OneOf: 2011-08-15 18:18:18:1;1970-01-01 00:00:00:1
float1:
  # Generates random floats uniformly at random from [MinVal, MaxVal).
  Type: float  # Alias for `float/uniform`.
  MinVal: -0.1  # Minimal value.
  MaxVal: 10.9  # Maximal value.
float2:
  # Generates floats from [-math.MaxFloat64, math.MaxFloat64] with normal distribution
  # `mean=(MinVal+MaxVal)/2=20` and `stDev=(MaxVal-MinVal)/2=10`.
  Type: float/normal
  MinVal: 10  # 16% of values will be smaller than this (mean - stDev).
  MaxVal: 30  # 16% of values will be greater than this (mean + stDev).
float3:
  # Generates floats from (MinVal,  math.MaxFloat64] with exponential distribution
  # `mean=(MinVal+MaxVal)/2` and ~15% of values are greater than MaxVal.
  Type: float/exp
  MinVal: 10  # All values will be greater than this.
  MaxVal: 30  # 15% of values will be greater than this.
float4:
  # Selects one of the numbers given in OneOf.
  Type: float/oneof
  OneOf: 0.5;1.5;6.5:10  # `0.5`, `1.5`, or `6.5` with the latter being 10 times more likely.
int1:
  # Generates random ints uniformly at random from [MinVal, MaxVal).
  Type: int  # Alias for `int/uniform`.
  MinVal: -10  # Minimal value.
  MaxVal: 100  # Maximal value.
int2:
  # Selects one of the numbers given in OneOf.
  Type: int/oneof
  OneOf: 0;1;6:10  # `0`, `1`, or `6` with the latter being 10 times more likely.
string1:
  # Generates random strings of given length.
  Type: string  # Alias for `string/rand`.
  Length: 15  # Strings will be 15 characters long.
  # [Optional] Provide character set to pick from.
  OneOf: "abc xyz"  # Default is letters (upper/lower case) and numbers.
string2:
  # Generates random sections of lorem ipsum text of given length.
  Type: string/text
  Length: 150  # Number of characters in the output string.
string3:
  # Generates a 36 characters long universally unique string.
  Type: string/uuid
string4:
  # Selects one of the options given in OneOf.
  Type: string/oneof
  OneOf: yes:95;no:5  # Generates `yes` 95% of the time and `no` 5% of the time.
string5:
  # Any type can be partly `NULL` by setting the nullable field.
  Type: string/uuid
  Nullable: 0.3  # This will be `NULL` 30% of the time and random UUID 70% of the time.
```
### Additional note on the OneOf field
Some extra information should be provided on the topic of the `OneOf` field. In the `*/oneof` generator types this field
is used to list all the possible values and their respective weights. Different options are delimited with semicolons
whereas an option and its weight are delimited with a colon.
```yaml
OneOf: "<option 1>:<weight 1>;<option 2>:<weight 2>;...;<option n>:<weight n>"
```
Weights have to be non-negative integers, and they have to sum to a positive value (i.e. at least one has to be
positive). The default value for weights is `1`, so you don't have to write them if you don't want to change them. The
only exception being if the option itself includes one or more colons in which case you have to manually append `:1` to
the end. There is currently no way for an option to include a literal `;`. See examples below:

```yaml
# All options have weight 1, so they will be picked uniformly at random.
OneOf: "<option 1>;<option 2>;<option 3>"
```
```yaml
# Options 1 and 3 have weight 1 and option 2 has weight 10 making it 10 times more likely to be picked.
OneOf: "<option 1>;<option 2>:10;<option 3>"
```

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