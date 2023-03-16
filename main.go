package main

import (
	"bufio"
	"command-to-vsconfig/parser"
	"os"
	"strings"

	"encoding/json"

	"github.com/antlr/antlr4/runtime/Go/antlr/v4"
)

func getType(binName string) string {
	switch binName {
	case "python":
		return "python"
	case "python3":
		return "python"
	case "gcc":
		return "cpp"
	default:
		return "cpp"
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

	for _, arg := range tree.AllArg() {
		println(arg.GetText())
	}

	argsList := []string{}
	for _, item := range tree.AllArg() {
		argsList = append(argsList, item.GetText())
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
