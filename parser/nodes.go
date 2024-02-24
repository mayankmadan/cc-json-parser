package parser

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
