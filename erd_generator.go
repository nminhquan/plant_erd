package main

import (
	"fmt"
	"github.com/sue445/plant_erd/db"
	"io/ioutil"
	"os"
)

// ErdGenerator represents ERD generator
type ErdGenerator struct {
	Filepath string
	Table    string
	Distance int
}

// Run performs generator
func (g *ErdGenerator) Run(schema *db.Schema) error {
	err := g.checkParamTable(schema)

	if err != nil {
		return err
	}

	erd := g.generate(schema)
	return g.outputErd(erd)
}

func (g *ErdGenerator) checkParamTable(schema *db.Schema) error {
	if g.Table == "" {
		return nil
	}

	for _, table := range schema.Tables {
		if table.Name == g.Table {
			return nil
		}
	}

	return fmt.Errorf("%s is not found in database", g.Table)
}

func (g *ErdGenerator) generate(schema *db.Schema) string {
	if g.Table == "" || g.Distance <= 0 {
		return schema.ToErd()
	}

	subset := schema.Subset(g.Table, g.Distance)
	return subset.ToErd()
}

func (g *ErdGenerator) outputErd(content string) error {
	if g.Filepath == "" {
		// Print to stdout
		fmt.Fprint(os.Stdout, content)
		return nil
	}

	// Output to file
	return ioutil.WriteFile(g.Filepath, []byte(content), 0644)
}
