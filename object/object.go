package object

import (
	"fmt"
)

type ObjectType string

const (
	INTEGER_OBF = "INTEGER"
	BOOLEAN_OBJ = "BOOLEAN"
)

type object interface {
	Type() ObjectType 
	Inspect() string
}


type Boolean struct {
	Value bool
}

func (b *Boolean) Type() ObjectType { return BOOLEAN_OBJ}
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t",b.Value)}