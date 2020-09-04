package main

import (
	"fmt"
	"github.com/hagarihayato/mercari2020/usecase"
	"go/ast"
	"go/parser"
	"go/token"
	"syscall/js"
)

//var array = []string{"fmt", "go/ast", "strings", "golang.org/x/tools/go/packages"}
//var expr = "//*[@type='CallExpr']/Fun[@type='Ident' and @Name='panic']"
var document = js.Global().Get("document")
var importFile = document.Call("getElementById", "file")
var condition = document.Call("getElementById", "condition")
var wrapper = document.Call("getElementById", "wrapper")


func main() {
	c := make(chan struct{}, 0)
	println("Go WebAssembly Initialized")
	registerCallbacks()
	<-c
	select {}
}

func registerCallbacks() {
	js.Global().Set("pushBtn", js.FuncOf(pushBtn))
}

func pushBtn(this js.Value, args []js.Value) interface{} {
	//expr := condition.Get("value").String()
	expr := "//*[@type=\"CallExpr\"]/Fun[@type=\"Ident\" and @Name=\"len\"]"
	fileContent := wrapper.Get("value")
	fs := token.NewFileSet()
	f, _ := parser.ParseFile(fs, "main.go", fileContent, 0)
	fmt.Println(fileContent)
	ast.Print(fs, f)
	//// astquery実行
	query, err := usecase.QueryLoader(fs, expr, f)
	fmt.Println(query)
	if err != nil {
		panic(err)
	}

	// フォーム初期化
	//importFile.Set("files[0]", "")
	//importFile.Set("value", "")
	condition.Set("value", "")
	return nil
}

