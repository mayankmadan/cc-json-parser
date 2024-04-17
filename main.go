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

func printTree(nodes []parser.JsonNode, level int) {
	if len(nodes) == 0 {
		return
	}
	for _, node := range nodes {
		for i := 0; i < level; i++ {
			fmt.Printf("\t")
		}
		fmt.Printf("type: %d, key: %s, value: %s\n", node.Type(), node.Key(), node.Value())
		// if node.Type() == parser.NodeTypeArray {
		// 	fmt.Printf("Array: %v\n", node)
		// }
		printTree(node.Children(), level+1)
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
	node, err := p.Parse()

	if err != nil {
		log.Println(err.Error())
		return
	}
	node.Put("hello", parser.NewNode("World!"))

	printTree([]parser.JsonNode{node}, 0)
	newNode := parser.NewNode("This is a put test")
	newNode2 := parser.NewNode(123)
	newNode3 := parser.NewNode(true)
	newNode4 := parser.NewNode(nil)
	newNode5 := parser.NewNode([]any{1, 2, 3, 100})
	node.Get("test5").Get("a").Put("putTest", newNode)
	node.Get("test5").Get("a").Put("putTest2", newNode2)
	node.Get("test5").Get("a").Put("putTest3", newNode3)
	node.Get("test5").Get("a").Put("putTest4", newNode4)
	node.Put("arrayNode", newNode5)
	printTree([]parser.JsonNode{node}, 0)
}
