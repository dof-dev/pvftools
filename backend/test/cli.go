package main

import (
	"encoding/json"
	"os"
	"pvftools/backend/internal/breath"

	"github.com/dof-dev/godof/parser/tree_parser"
)

func do(path, ret string) {
	c, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	parser := tree_parser.NewTreeParser(string(c))
	e := breath.NewBreathEditor(1254, parser)
	buf, _ := json.Marshal(e.Data)
	os.WriteFile(ret, buf, 0644)
}

func main() {
	do("/Users/ziipin/Projects/dof/pvf/stackable/consumption_1254.stk", "/tmp/1.json")
	//do("/tmp/1.stk", "/tmp/2.json")

	//ret := e.Render()
	//os.WriteFile("/tmp/1.stk", []byte(ret), 0644)
}
