package object

import "fmt"

type ObjectType string

const (
	OBJ_TYPE_INTEGER = "Integer"
	OBJ_TYPE_BOOLEAN = "Boolean"
	OBJ_TYPE_NULL    = "Null"
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
