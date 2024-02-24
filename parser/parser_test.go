package parser

import (
	"reflect"
	"testing"

	"github.com/mayankmadan/jsonparser/lexer"
)

func TestParser_generateAST(t *testing.T) {
	type args struct {
		data string
	}
	tests := []struct {
		name    string
		args    args
		want    *AST
		wantErr bool
	}{
		{
			name: "Empty",
			args: args{
				data: "",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Valid JSON",
			args: args{
				data: `{"key": "value"}`,
			},
			want:    &AST{Type: NodeTypeObject, Key: "", Value: nil, Children: []*AST{{Type: NodeTypeString, Key: "key", Value: "value"}}},
			wantErr: false,
		},
		{
			name: "Invalid JSON",
			args: args{
				data: `{"key": "value"`,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Nested JSON",
			args: args{
				data: `{"key1": "value", "key2": {"key3": "value3"}, "test4": "value4"}`,
			},
			want:    &AST{Type: NodeTypeObject, Key: "", Value: nil, Children: []*AST{{Type: NodeTypeString, Key: "key1", Value: "value", Children: nil}, {Type: NodeTypeObject, Key: "key2", Value: nil, Children: []*AST{{Type: NodeTypeString, Key: "key3", Value: "value3"}}}, {Type: NodeTypeString, Key: "test4", Value: "value4"}}},
			wantErr: false,
		},
		{
			name: "Trailing Comma Invalid",
			args: args{
				data: `{"key1": "value", "key2": {"key3": "value3"}, "test4": "value4",}`,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Nested Array",
			args: args{
				data: `{"key1": "value", "key2": {"key3": ["value3", "value4"]}, "test4": "value4"}`,
			},
			want:    &AST{Type: NodeTypeObject, Key: "", Value: nil, Children: []*AST{{Type: NodeTypeString, Key: "key1", Value: "value", Children: nil}, {Type: NodeTypeObject, Key: "key2", Value: nil, Children: []*AST{{Type: NodeTypeArray, Key: "key3", Value: nil, Children: []*AST{{Type: NodeTypeString, Key: "", Value: "value3"}, {Type: NodeTypeString, Key: "", Value: "value4"}}}}}, {Type: NodeTypeString, Key: "test4", Value: "value4"}}},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := lexer.NewLexer([]rune(tt.args.data))
			err := l.Lex()
			if err != nil {
				t.Errorf("Lexer.Lex() error = %v", err)
				return
			}
			p := NewParser(l.Tokens())
			got, err := p.Parse()
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}
