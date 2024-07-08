package object

import (
	"bytes"
	"fmt"
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
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string  { return fmt.Sprintf("%d", i.Value) }
func (i *Integer) Type() ObjectType { return OBJ_TYPE_INTEGER }

type Boolean struct {
	Value bool
}

func (b *Boolean) Inspect() string  { return fmt.Sprintf("%t", b.Value) }
func (b *Boolean) Type() ObjectType { return OBJ_TYPE_BOOLEAN }

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
	Value string
}

func (s *String) Type() ObjectType { return OBJ_TYPE_STRING }
func (s *String) Inspect() string  { return s.Value }

type BuiltinFunction func(args ...Object) Object

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType { return OBJ_TYPE_BUILTIN }
func (b *Builtin) Inspect() string  { return "builtin function" }
