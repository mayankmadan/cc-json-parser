package main

import (
	"fmt"
	"log"
	"os"

	"github.com/mayankmadan/jsonparser/lexer"
	"github.com/mayankmadan/jsonparser/parser"
)

func printTokens(tokens []lexer.Token) {
	for _, v := range tokens {
		fmt.Printf("%v\n", v)
	}
}

func printTree(nodes []*parser.AST, level int) {
	if len(nodes) == 0 {
		return
	}
	for _, node := range nodes {
		for i := 0; i < level; i++ {
			fmt.Printf("\t")
		}
		fmt.Printf("type: %d, key: %s, value: %s\n", node.Type, node.Key, node.Value)
		printTree(node.Children, level+1)
	}
}

func main() {
	log.SetFlags(log.Lshortfile)
	data, err := os.ReadFile("test/test2.json")
	if err != nil {
		panic(err)
	}
	input := []rune(string(data))
	lexer := lexer.NewLexer(input)
	err = lexer.Lex()
	if err != nil {
		fmt.Printf("%v\n", err)
		return
	}
	p := parser.NewParser(lexer.Tokens())
	tree, err := p.Parse()

	if err != nil {
		log.Println(err.Error())
		return
	}

	printTree([]*parser.AST{tree}, 0)
}
