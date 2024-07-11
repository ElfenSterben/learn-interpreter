package object

import (
	"bytes"
	"fmt"
	"hash/fnv"
	"learn-interpreter/ast"
	"strings"
)

type ObjectType string

const (
	OBJ_TYPE_INTEGER      = "Integer"
	OBJ_TYPE_BOOLEAN      = "Boolean"
	OBJ_TYPE_NULL         = "Null"
	OBJ_TYPE_RETURN_VALUE = "ReturnValue"
	OBJ_TYPE_ERROR        = "Error"
	OBJ_TYPE_FUNCTION     = "Function"
	OBJ_TYPE_STRING       = "String"
	OBJ_TYPE_BUILTIN      = "BuiltIn"
	OBJ_TYPE_ARRAY        = "Array"
	OBJ_TYPE_HASH         = "Hash"
	OBJ_TYPE_QUOTE        = "Quote"
	OBJ_TYPE_MACRO        = "Macro"
)

type HashKey struct {
	Type  ObjectType
	Value uint64
}

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Hashable interface {
	HashKey() HashKey
}

type Integer struct {
	Value   int64
	hashKey HashKey
}

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return OBJ_TYPE_INTEGER }
func (i *Integer) HashKey() HashKey {
	if i.hashKey.Value == 0 && i.hashKey.Type == "" {
		i.hashKey = HashKey{Type: i.Type(), Value: uint64(i.Value)}
	}
	return i.hashKey
}

type Boolean struct {
	Value   bool
	hashKey HashKey
}

func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return OBJ_TYPE_BOOLEAN }
func (b *Boolean) HashKey() HashKey {
	if b.hashKey.Value == 0 && b.hashKey.Type == "" {
		var value uint64
		if b.Value {
			value = 1
		} else {
			value = 0
		}
		b.hashKey = HashKey{Type: b.Type(), Value: value}
	}
	return b.hashKey
}

type Null struct{}

func (n *Null) Inspect() string  { return "null" }
func (n *Null) Type() ObjectType { return OBJ_TYPE_NULL }

type ReturnValue struct {
	Value Object
}

func (rv *ReturnValue) Inspect() string  { return rv.Value.Inspect() }
func (rv *ReturnValue) Type() ObjectType { return OBJ_TYPE_RETURN_VALUE }

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType { return OBJ_TYPE_ERROR }
func (e *Error) Inspect() string  { return "ERROR: " + e.Message }

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() ObjectType { return OBJ_TYPE_FUNCTION }
func (f *Function) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")
	return out.String()
}

type String struct {
	Value   string
	hashKey HashKey
}

func (s *String) Type() ObjectType { return OBJ_TYPE_STRING }
func (s *String) Inspect() string  { return s.Value }
func (s *String) HashKey() HashKey {
	if s.hashKey.Value == 0 && s.hashKey.Type == "" {
		h := fnv.New64a()
		h.Write([]byte(s.Value))
		s.hashKey = HashKey{Type: s.Type(), Value: h.Sum64()}
	}
	return s.hashKey
}

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return OBJ_TYPE_BUILTIN }
func (b *Builtin) Inspect() string  { return "builtin function" }

type Array struct {
	Elements []Object
}

func (ao *Array) Type() ObjectType { return OBJ_TYPE_ARRAY }
func (ao *Array) Inspect() string {
	var out bytes.Buffer
	elements := []string{}
	for _, e := range ao.Elements {
		elements = append(elements, e.Inspect())
	}
	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")
	return out.String()
}

type HashPair struct {
	Key   Object
	Value Object
}

type Hash struct {
	Pairs map[HashKey]HashPair
}

func (h *Hash) Type() ObjectType { return OBJ_TYPE_HASH }
func (h *Hash) Inspect() string {
	var out bytes.Buffer
	pairs := []string{}
	for _, pair := range h.Pairs {
		pairs = append(pairs, fmt.Sprintf("%s: %s",
			pair.Key.Inspect(), pair.Value.Inspect()))
	}
	out.WriteString("{")
	out.WriteString(strings.Join(pairs, ", "))
	out.WriteString("}")
	return out.String()
}

type Quote struct {
	Node ast.Node
}

func (q *Quote) Type() ObjectType { return OBJ_TYPE_QUOTE }
func (q *Quote) Inspect() string  { return "QUOTE(" + q.Node.String() + ")" }

type Macro struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (m *Macro) Type() ObjectType { return OBJ_TYPE_MACRO }
func (m *Macro) Inspect() string {
	var out bytes.Buffer
	params := []string{}
	for _, p := range m.Parameters {
		params = append(params, p.String())
	}
	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(m.Body.String())
	out.WriteString("\n}")
	return out.String()
}
