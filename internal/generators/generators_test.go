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

type readmeConfig map[string]config.Args

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
		"example1":  "1",
		"example2":  "1",
		"example3":  "NOW()",
		"example4":  "'2016-12-19 23:42:51'",
		"example5":  "'1970-01-01 00:00:00'",
		"example6":  "2.577089127583518",
		"example7":  "27.133352143941206",
		"example8":  "12.181511239890659",
		"example9":  "1.5",
		"example10": "89",
		"example11": "1",
		"example12": "'bbbbyaayzcbzazx'",
		"example13": "'ibulum in. Fusce lacinia, mi vel viverra viverra, lacus velit vulputate justo, nec vehicula ipsum enim et ligula. Sed sed convallis ex. Nam lobortis a'",
		"example14": "'4d618232-ae05-46d0-a270-2931ef3d9add'",
		"example15": "'yes'",
		"example16": "NULL",
	}
	for col, expected := range tests {
		g, err := GetGenerator(c[col])
		if err != nil {
			t.Errorf("unable to load generator for %s (Type: %s): %s", col, c[col].Type, err)
		}
		assert.Equal(t, expected, g.Next(), "column %s (Type: %s) is incorrect", col, c[col].Type)
	}
}
