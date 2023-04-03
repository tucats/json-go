package main

import (
	"fmt"
	"os"
)

const helpText = `json-go: Generate Go data definitions from sample json

Usage: 
    json-go [options]

Options:
  --camel,  -c                   Use camel-case for struct member names
  --file,   -f    <filename>     Read sample json from named file instead of stdin
  --type,   -t    <typename>     Specify base type name in generated text
  --output, -o    <filename>     Write generated text to named file instead of stdout`

func help() {
	fmt.Println(helpText)
	os.Exit(0)
}
