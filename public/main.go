package main

import (
	"fmt"
	"github.com/gostaticanalysis/astquery"
	"go/ast"
	"golang.org/x/tools/go/packages"
	"os"
	//"syscall/js"
)

func main() {
	// グローバルオブジェクト（Webブラウザはwindow）の取得
	//window := js.Global()

	// window.document.getElementById("message")を実行
	//condition := window.Get("document").Call("getElementById", "condition").String()
	//packName := window.Get("document").Call("getElementById", "condition").String()
	//select{}
	array := []string{"fmt", "go/ast", "golang.org/x/tools/go/packages"}
	expr := "//*[@type='CallExpr']/Fun[@type='Ident' and @Name='panic']"
	query := queryLoader(expr, array)
	fmt.Println("------------", query, "-------------")
}

func queryLoader(expr string, array []string) []string {
	cfg := &packages.Config{Mode: packages.NeedFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedDeps}
	pkgs, err := packages.Load(cfg, array...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "load: %v\n", err)
		os.Exit(1)
	}
	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1)
	}
	for _, pkg := range pkgs {
		e := astquery.New(pkg.Fset, pkg.Syntax, nil)
		v, err := e.Eval(expr)
		if err != nil {
			panic(err)
		}
		switch v := v.(type) {
		case []ast.Node:
			queryArray := []string{}
			for _, n := range v {
				fmt.Printf("%[1]T %[1]v\n", n)
				queryArray = append(queryArray, fmt.Sprintf("%[1]T %[1]v\n", n))
			}
			return queryArray
		default:
			//rv := reflect.ValueOf(v)
			//switch rv.Kind() {
			//case reflect.Array, reflect.Slice:
			//	for i := 0; i < rv.Len(); i++ {
			//		fmt.Println(rv.Index(i).Interface())
			//	}
			//case reflect.Map:
			//	for _, key := range rv.MapKeys() {
			//		val := rv.MapIndex(key)
			//		fmt.Printf("%v:%v\n", key.Interface(), val.Interface())
			//	}
			//default:
			//	fmt.Println(v)
			//}
		}
	}
	return nil
}

