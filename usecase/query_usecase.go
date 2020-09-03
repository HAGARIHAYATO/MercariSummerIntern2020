package usecase

import (
	"fmt"
	"github.com/gostaticanalysis/astquery"
	"go/ast"
	"golang.org/x/tools/go/packages"
)

// wasm挙動確認用関数
func QueryLoad(expr string, array ...string) ([]string, error) {
	var list []string
	for _, item := range array {
		s := fmt.Sprintf("%[1]v %[2]v", expr, item)
		list = append(list, s)
	}
	return list, nil
}

func QueryLoader(expr string, array ...string) ([]string, error) {
	cfg := &packages.Config{Mode: packages.NeedFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedDeps}
	// 時間がかかるからwasmの処理で弾かれる？
	pList, err := packages.Load(cfg, array...)
	if err != nil {
		return nil, err
	}
	var queryArray []string
	for _, pkg := range pList {
		e := astquery.New(pkg.Fset, pkg.Syntax, nil)
		v, err := e.Eval(expr)
		if err != nil {
			return nil, err
		}
		switch v := v.(type) {
		case []ast.Node:
			for _, n := range v {
				queryArray = append(queryArray, fmt.Sprintf("%[1]T %[1]v\n", n))
			}
		}
	}
	return queryArray, nil
}
