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
var hiddenField = document.Call("getElementById", "hiddenField")
var prefix = document.Call("getElementById", "prefix")
var terminal = document.Call("getElementById", "terminal")
var message = document.Call("getElementById", "message")

func main() {
	c := make(chan struct{}, 0)
	println("Go WebAssembly Initialized")
	registerCallbacks()
	<-c
	select {}
}

func registerCallbacks() {
	js.Global().Set("pushBtn", js.FuncOf(pushBtn))
	js.Global().Set("resetBtn", js.FuncOf(resetBtn))
}

func pushBtn(this js.Value, args []js.Value) interface{} {
	expr := condition.Get("value").String()
	fileContent := hiddenField.Get("value").String()

	if expr == "" { return nil }
	if fileContent == "" { return nil }

	// ターミナル内の処理
	if prefix.Get("innerText").String() == "~ $" {
		prefix.Set("innerText", "~ $ astquery" + "  " + "'" + expr + "'" + "  " + importFile.Get("value").String())
	} else {
		pre := document.Call("createElement", "p")
		pre.Set("innerText", "~ $ astquery" + "  " + "'" + expr + "'" + "  " + importFile.Get("value").String())
		terminal.Call("appendChild", pre)
	}

	fs := token.NewFileSet()
	f, _ := parser.ParseFile(fs, "main.go", fileContent, 0)


	ast.Print(fs, f)


	//// astquery実行
	query, err := usecase.QueryLoader(fs, expr, f)
	if err != nil {
		message.Set("innerText", fmt.Sprintf("error: %[1]s , please reload this page", err))
		panic(err)
	}

	// 帰ってきたクエリを展開
	for _, q := range query {
		paragraph := document.Call("createElement", "p")
		paragraph.Set("innerText", fmt.Sprintf("%[1]T %[1]v\n", q))
		terminal.Call("appendChild", paragraph)
	}

	return nil
}

func resetBtn(this js.Value, args []js.Value) interface{} {
	// フォーム初期化
	terminal.Set("innerHTML", "")
	prefix.Set("innerText", "~ $")
	return nil
}

