package parser

import (
	"errors"
	"fmt"

	"github.com/mayankmadan/jsonparser/lexer"
)

type AST struct {
	Type     NodeType
	Key      string
	Value    any
	Children []*AST
}

type Parser struct {
	tokens []lexer.Token
	pos    int
	root   *AST
}

func NewParser(tokens []lexer.Token) *Parser {
	return &Parser{tokens: tokens}
}

func (p *Parser) getToken() (*lexer.Token, error) {
	token, err := p.peekToken()
	if err != nil {
		return nil, err
	}
	p.pos++
	return token, nil
}

func (p *Parser) peekToken() (*lexer.Token, error) {
	if p.pos >= len(p.tokens) {
		return nil, errors.New("unexpected end of input")
	}
	token := p.tokens[p.pos]
	return &token, nil
}

func (p *Parser) getNodeType(currentToken lexer.Token) NodeType {

	switch currentToken.Type {
	case lexer.LBRACKET:
		return NodeTypeArray
	case lexer.LBRACE:
		return NodeTypeObject
	case lexer.STRING:
		return NodeTypeString
	case lexer.NUMBER:
		return NodeTypeNumber
	case lexer.BOOLEAN:
		return NodeTypeBoolean
	default:
		return NodeTypeInvalid
	}
}

func (p *Parser) getKey() (string, error) {
	currentToken, err := p.getToken()
	if err != nil {
		return "", err
	}
	if currentToken.Type != lexer.STRING {
		return "", errors.New("unexpected token " + currentToken.Value + " at position: " + fmt.Sprint(currentToken.Pos))
	}

	colon, err := p.getToken()
	if err != nil {
		return "", err
	}

	if colon.Type != lexer.COLON {
		return "", errors.New("unexpected token " + currentToken.Value + " at position: " + fmt.Sprint(currentToken.Pos))
	}
	return currentToken.Value, nil
}

func (p *Parser) getValue() (*AST, error) {
	currentToken, err := p.peekToken()
	if err != nil {
		return nil, err
	}
	nodeType := p.getNodeType(*currentToken)
	node := AST{Type: nodeType}

	switch nodeType {
	case NodeTypeObject:
		node.Children, err = p.parseObject()
		if err != nil {
			return nil, err
		}
	case NodeTypeArray:
		node.Children, err = p.parseArray()
		if err != nil {
			return nil, err
		}

	case NodeTypeString, NodeTypeNumber, NodeTypeBoolean, NodeTypeNull:
		node.Value = currentToken.Value
	default:
		return nil, fmt.Errorf("unexpected token %s at position: %d", currentToken.Value, currentToken.Pos)
	}
	p.pos++
	return &node, nil
}

func (p *Parser) parseArray() ([]*AST, error) {
	children := []*AST{}
	p.pos++
	trailingComma := true

	for {
		currentToken, err := p.peekToken()
		if err != nil {
			return nil, err
		}
		if currentToken.Type == lexer.RBRACKET {
			if trailingComma && len(children) > 0 {
				return nil, fmt.Errorf("unexpected token %s at position: %d", currentToken.Value, currentToken.Pos)
			}
			break
		}
		trailingComma = false

		node, err := p.getValue()
		if err != nil {
			return nil, err
		}
		children = append(children, node)
		currentToken, err = p.peekToken()
		if err != nil {
			return nil, err
		}
		if currentToken.Type == lexer.COMMA {
			trailingComma = true
			p.pos++
		}
	}

	return children, nil
}

func (p *Parser) parseObject() ([]*AST, error) {
	children := []*AST{}
	if p.pos >= len(p.tokens) {
		return nil, errors.New("unexpected end of input")
	}
	p.pos++
	trailingComma := true
	for {
		currentToken, err := p.peekToken()
		if err != nil {
			return nil, err
		}
		if currentToken.Type == lexer.RBRACE {
			if trailingComma && len(children) > 0 {
				return nil, fmt.Errorf("unexpected token %s at position: %d", currentToken.Value, currentToken.Pos)
			}
			break
		}
		if !trailingComma {
			return nil, fmt.Errorf("unexpected token %s at position: %d", currentToken.Value, currentToken.Pos)
		}
		trailingComma = false
		key, err := p.getKey()
		if err != nil {
			return nil, err
		}

		node, err := p.getValue()
		if err != nil {
			return nil, err
		}
		node.Key = key
		children = append(children, node)

		currentToken, err = p.peekToken()
		if err != nil {
			return nil, err
		}

		if currentToken.Type == lexer.COMMA {
			trailingComma = true
			p.pos++
		}
	}

	return children, nil
}

func (p *Parser) Parse() (*AST, error) {
	var err error
	if len(p.tokens) == 0 {
		return nil, errors.New("input empty")
	}
	p.pos = 0
	currentToken, err := p.peekToken()
	if err != nil {
		return nil, err
	}
	rootNodeType := p.getNodeType(*currentToken)
	if rootNodeType != NodeTypeArray && rootNodeType != NodeTypeObject {
		return nil, errors.New("unexpected token " + currentToken.Value + " at position: " + fmt.Sprint(currentToken.Pos))
	}
	var rootNode *AST

	switch rootNodeType {
	case NodeTypeArray:
		rootNode = &AST{Type: NodeTypeArray}
		rootNode.Children, err = p.parseArray()
		if err != nil {
			return nil, err
		}
		p.root = rootNode
		return rootNode, nil
	case NodeTypeObject:
		rootNode = &AST{Type: NodeTypeObject}

		rootNode.Children, err = p.parseObject()
		if err != nil {
			return nil, err
		}
		p.root = rootNode
		return rootNode, nil
	}
	return nil, errors.New("root node needs to be either array or object")
}
