# cc-json-parser

A simple jsonparser for codingchallenges.fyi (https://codingchallenges.fyi/challenges/challenge-json-parser/)

This project provides a Go library for parsing JSON data. It includes a lexer, parser, and utilities for manipulating JSON nodes.

### Features

* Lexical analysis (tokenization) of JSON data
* Parsing of JSON data into a tree structure
* Representation of JSON data using nodes with key, value, type, and children
* Basic manipulation of JSON nodes (get and put operations)

### How to Use

1. **Install the library:**

```bash
go get github.com/mayankmadan/cc-json-parser
```

2. **Import the library:**

```go
import (
"github.com/mayankmadan/cc-json-parser/lexer"
"github.com/mayankmadan/cc-json-parser/parser"
)
```

3. **Lex and parse JSON data:**

```go
// Example JSON data
data := []rune(`{"name": "John Doe", "age": 30}`)

// Create a lexer
lexer := lexer.NewLexer(data)
err := lexer.Lex()
if err != nil {
// Handle error
}

// Create a parser
p := parser.NewParser(lexer.Tokens())

// Parse the data
node, err := p.Parse()
if err != nil {
// Handle error
}
```

4. **Manipulate JSON nodes:**

```go
// Get the "name" field
nameNode := node.Get("name")

// Change the value of the "name" field
nameNode.Value = "Jane Doe"

// Add a new field
node.Put("occupation", parser.NewNode("Software Engineer"))
```

### Understanding the Code

* **lexer/lexer.go:** Implements the lexical analysis, transforming the raw JSON data into tokens (e.g., LBRACE, RBRACE, STRING, NUMBER).
* **parser/parser.go:** Handles the parsing of the token stream into a tree structure represented by `JsonNode` instances.
* **parser/ast.go:** Defines the `AST` struct, which implements the `JsonNode` interface and stores the key, value, type, and children of each node.
* **parser/nodes.go:** Defines the `JsonNode` interface and provides a constructor function `NewNode` to create nodes of different types.

### Limitations

* This parser is a basic implementation and may not handle all edge cases or complex JSON structures.
* Error handling is rudimentary and could be improved.
* Functionality is limited to basic get and put operations.

### Potential Improvements

* Enhance error handling with more informative messages and error types.
* Implement additional manipulation methods like delete, update, and search.
* Support more complex JSON features like comments and escaped characters.
* Add unit tests for better code coverage and reliability.

### Contributions

Feel free to fork the repository and contribute improvements or bug fixes. Pull requests are welcome!