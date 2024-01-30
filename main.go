package main

import (
	"encoding/json"
	"os"
)

func main() {
	args := os.Args
	if len(args) > 1 {
		if args[1] == "supports" {
			os.Exit(0)
		}
	}

	var mdbook []map[string]interface{}
	err := json.NewDecoder(os.Stdin).Decode(&mdbook)
	if err != nil {
		panic(err)
	}
	if len(mdbook) < 2 {
		os.Exit(1)
	}
	context := mdbook[0]
	book := mdbook[1]

}
