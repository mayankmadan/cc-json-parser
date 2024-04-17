package parser

type AST struct {
	nodeType NodeType
	key      string
	value    any
	children []JsonNode
}

func (a *AST) Type() NodeType {
	return a.nodeType
}

func (a *AST) Key() string {
	return a.key
}

func (a *AST) Value() any {
	return a.value
}

func (a *AST) Children() []JsonNode {
	return a.children
}

func (a *AST) Get(key string) JsonNode {
	if a.nodeType != NodeTypeObject {
		return nil
	}
	for _, child := range a.children {
		if child.Key() == key {
			return child
		}
	}
	return nil
}

func (a *AST) GetByIndex(index int) JsonNode {
	if a.nodeType != NodeTypeArray {
		return nil
	}
	if index >= 0 && index < len(a.children) {
		return a.children[index]
	}
	return nil
}

func (a *AST) Put(key string, val JsonNode) bool {
	if a.nodeType != NodeTypeObject {
		return false
	}
	if key == "" {
		return false
	}
	node := &AST{nodeType: val.Type(), key: key, value: val.Value()}
	a.children = append(a.children, node)
	return true
}
