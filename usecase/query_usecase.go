package usecase

import (
	"fmt"
	"github.com/gostaticanalysis/astquery"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

// wasm挙動確認用関数
func QueryLoad(expr string, array []string) ([]string, error) {
	var list []string
	for _, item := range array {
		s := fmt.Sprintf("%[1]v %[2]v", expr, item)
		list = append(list, s)
	}
	return list, nil
}

func QueryLoader(expr string, array []string) ([]ast.Node, error) {
	cfg := &packages.Config{Mode: packages.NeedFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedDeps}
	// 時間がかかるからwasmの処理で弾かれる？
	// error ex... "'go list' driver requires 'go', but executable file not found in $PATH"
	pList, err := packages.Load(cfg, array...)
	fmt.Println(pList, err)
	if err != nil {
		return nil, err
	}
	var queryArray []ast.Node
	for _, pkg := range pList {
		e := astquery.New(pkg.Fset, pkg.Syntax, nil)
		v, err := e.Eval(expr)
		if err != nil {
			return nil, err
		}
		switch v := v.(type) {
		case []ast.Node:
			queryArray = v
		}
	}
	return queryArray, nil
}
