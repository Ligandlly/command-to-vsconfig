package main

import (
	"bufio"
	"os"
	"strings"

	"github.com/Ligandlly/command-to-vsconfig/parser"

	"encoding/json"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func getType(binName string) string {
	switch binName {
	case "python":
		return "python"
	case "python3":
		return "python"
	case "g++":
		return "cpp"
	default:
		return "python"
	}
}

func main() {
	println("Enter command: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := strings.TrimSpace(scanner.Text())
	inputStream := antlr.NewInputStream(input)

	lexer := parser.NewCommandLexer(inputStream)
	stream := antlr.NewCommonTokenStream(lexer, 0)
	p := parser.NewCommandParser(stream)
	p.BuildParseTrees = true
	tree := p.Command()

	argsList := []string{}
	for _, item := range tree.AllArg() {
		key := item.Key().GetText()
		if item.Value() != nil {
			value := item.Value().GetText()
			argsList = append(argsList, key, value)
		} else {
			argsList = append(argsList, key)
		}
	}

	bin_name := tree.Bin().GetText()

	vsConfig := map[string]interface{}{
		"name":       "Python: Current File",
		"type":       getType(bin_name),
		"request":    "launch",
		"program":    tree.File().GetText(),
		"args":       argsList,
		"justMyCode": false,
	}

	println("VSCode config:")
	json.NewEncoder(os.Stdout).Encode(vsConfig)
}
