package usecase

import (
	"github.com/gostaticanalysis/astquery"
	"go/ast"
	"go/token"
)

// TODO
func QueryLoader(fs *token.FileSet, expr string, f *ast.File) ([]ast.Node, error) {
	var queryArray []ast.Node
	var ff = []*ast.File{f}
	e := astquery.New(fs, ff, nil)
	v, err := e.Eval(expr)
	if err != nil {
		return nil, err
	}
	switch v := v.(type) {
	case []ast.Node:
		queryArray = v
	}
	return queryArray, nil
}