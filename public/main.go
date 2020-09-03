package main

import (
	"fmt"
	"github.com/hagarihayato/mercari2020/usecase"
	"strings"
	"syscall/js"
)

//var array = []string{"fmt", "go/ast", "strings", "golang.org/x/tools/go/packages"}
//var expr = "//*[@type='CallExpr']/Fun[@type='Ident' and @Name='panic']"
var document = js.Global().Get("document")
var prefix = document.Call("getElementById", "prefix")
var terminal = document.Call("getElementById", "terminal")
var condition = document.Call("getElementById", "condition")
var pkg = document.Call("getElementById", "packName")



func main() {
	c := make(chan struct{}, 0)
	println("Go WebAssembly Initialized")
	registerCallbacks()
	<-c
}

func registerCallbacks() {
	js.Global().Set("pushBtn", js.FuncOf(pushBtn))
	js.Global().Set("resetBtn", js.FuncOf(resetBtn))
}

// terminal内をリセット
func resetBtn(this js.Value, args []js.Value) interface{} {
	terminal.Set("innerHTML", "")
	prefix.Set("innerText", "~ $")
	return nil
}

func pushBtn(this js.Value, args []js.Value) interface{} {
	expr := condition.Get("value").String()
	packName := pkg.Get("value").String()
	array := strings.Fields(packName)

	// terminal内のHTML書き換え
	if expr == "" || packName == "" { return nil }
	if prefix.Get("innerText").String() == "~ $" {
		prefix.Set("innerText", "~ $ astquery" + "  " + "'" + expr + "'" + "  " + packName)
	} else {
		pre := createElement("p")
		pre.Set("innerText", "~ $ astquery" + "  " + "'" + expr + "'" + "  " + packName)
		terminal.Call("appendChild", pre)
	}

	// astquery実行
	query, err := usecase.QueryLoader(expr, array...)
	if err != nil {
		fmt.Println(err)
	}

	// 帰ってきたクエリ配列をターミナルに展開
	for _, q := range query {
		paragraph := createElement("p")
		paragraph.Set("innerHTML", q)
		terminal.Call("appendChild", paragraph)
	}

	// フォーム初期化
	pkg.Set("value", "")
	condition.Set("value", "")
	return nil
}

func createElement(elementName string) js.Value {
	return document.Call("createElement", elementName)
}

