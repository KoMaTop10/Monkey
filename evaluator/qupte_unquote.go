package evaluator

import (
	"github.com/KoMaTop10/Monkey/ast"
	"github.com/KoMaTop10/Monkey/object"
)

func quote(node ast.Node) object.Object {
	return &object.Quote{Node: node}
}