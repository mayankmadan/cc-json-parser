package main

import (
	"fmt"
	"os"

	"github.com/mayankmadan/jsonparser/lexer"
)

func main() {
	data, err := os.ReadFile("test.json")
	if err != nil {
		panic(err)
	}
	input := []rune(string(data))
	lexer := lexer.NewLexer(input)
	err = lexer.Lex()
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	for _, v := range lexer.Tokens() {
		fmt.Printf("%v\n", v)
	}
}
