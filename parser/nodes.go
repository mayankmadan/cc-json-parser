package parser

import (
	"fmt"
	"reflect"
)

type NodeType int

const (
	NodeTypeInvalid NodeType = iota
	NodeTypeObject
	NodeTypeArray
	NodeTypeString
	NodeTypeNumber
	NodeTypeBoolean
	NodeTypeNull
)

type JsonNode interface {
	Key() string
	Value() any
	Children() []JsonNode
	Type() NodeType
	Get(key string) JsonNode
	GetByIndex(index int) JsonNode
	Put(key string, val JsonNode) bool
}

func NewNode(value any) JsonNode {
	if value == nil {
		return &AST{nodeType: NodeTypeNull, value: nil}
	}
	v := reflect.ValueOf(value)
	switch v.Kind() {
	case reflect.String:
		return &AST{nodeType: NodeTypeString, value: value}
	case reflect.Float32, reflect.Float64, reflect.Int, reflect.Int32, reflect.Int64, reflect.Int16, reflect.Int8:
		return &AST{nodeType: NodeTypeNumber, value: value}
	case reflect.Bool:
		return &AST{nodeType: NodeTypeBoolean, value: value}
	case reflect.Slice:
		children := []JsonNode{}
		for _, v := range value.([]any) {
			childNode := NewNode(v)
			children = append(children, childNode)
		}
		fmt.Printf("children: %v\n", children)

		return &AST{nodeType: NodeTypeArray, value: nil, children: children}
	default:
		return nil
	}
}
