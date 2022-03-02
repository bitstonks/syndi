package generators

import (
	"bufio"
	"fmt"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v2"
	"math/rand"
	"os"
	"testing"

	"github.com/bitstonks/syndi/internal/config"
)

func init() {
	// Override default RNG generator to create a deterministic one for tests
	newRng = func() *rand.Rand {
		rng := rand.New(rand.NewSource(4))
		return rng
	}
	uuidGen = func() string {
		return "4d618232-ae05-46d0-a270-2931ef3d9add"
	}
}

type readmeConfig map[string]config.ColumnDef

func loadConfigFromReadme() (readmeConfig, error) {
	var confData []byte
	file, err := os.Open("../../README.md")
	if err != nil {
		return nil, fmt.Errorf("unable to open README file: %s", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	isConfig := false
	for scanner.Scan() {
		line := scanner.Text()
		if isConfig && line == "```" {
			break
		}
		if line == "```yaml" {
			isConfig = true
			continue
		}
		if !isConfig {
			continue
		}
		confData = append(confData, []byte(line)...)
		confData = append(confData, byte('\n'))
	}
	c := readmeConfig{}
	err = yaml.Unmarshal(confData, &c)
	if err != nil {
		return nil, fmt.Errorf("unable to parse yaml from config: %s", err)
	}
	return c, err
}

func TestReadmeConfig(t *testing.T) {
	c, err := loadConfigFromReadme()
	if err != nil {
		t.Fatal(err)
	}
	tests := map[string]string{
		"bool1":     "1",
		"bool2":     "1",
		"datetime1": "NOW()",
		"datetime2": "'2016-12-19 23:42:51'",
		"datetime3": "'1970-01-01 00:00:00'",
		"float1":    "2.577089127583518",
		"float2":    "27.133352143941206",
		"float3":    "12.181511239890659",
		"float4":    "1.5",
		"int1":      "89",
		"int2":      "1",
		"string1":   "'bbbbyaayzcbzazx'",
		"string2":   "'ibulum in. Fusce lacinia, mi vel viverra viverra, lacus velit vulputate justo, nec vehicula ipsum enim et ligula. Sed sed convallis ex. Nam lobortis a'",
		"string3":   "'4d618232-ae05-46d0-a270-2931ef3d9add'",
		"string4":   "'yes'",
		"string5":   "NULL",
	}
	for col, _ := range c {
		if _, ok := tests[col]; !ok {
			t.Fatalf("config has an extra column: %s", col)
		}
	}
	for col, expected := range tests {
		conf, ok := c[col]
		if !ok {
			t.Fatalf("config is missing column %s", col)
		}
		g, err := GetGenerator(conf)
		if err != nil {
			t.Fatalf("unable to load generator for %s (Type: %s): %s", col, c[col].Type, err)
		}
		assert.Equal(t, expected, g.Next(), "column %s (Type: %s) is incorrect", col, c[col].Type)
	}
}
