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

func resetBtn(this js.Value, args []js.Value) interface{} {
	terminal.Set("innerHTML", "")
	prefix.Set("innerText", "~ $")
	return nil
}

func pushBtn(this js.Value, args []js.Value) interface{} {
	expr := condition.Get("value").String()
	packName := pkg.Get("value").String()
	array := strings.Fields(packName)
	if expr == "" || packName == "" { return nil }
	if prefix.Get("innerText").String() == "~ $" {
		prefix.Set("innerText", "~ $ astquery" + "  " + "'" + expr + "'" + "  " + packName)
	} else {
		pre := createElement("p")
		pre.Set("innerText", "~ $ astquery" + "  " + "'" + expr + "'" + "  " + packName)
		terminal.Call("appendChild", pre)
	}

	query, err := usecase.QueryLoader(expr, array...)
	if err != nil {
		fmt.Println(err)
	}

	for _, a := range query {
		paragraph := createElement("p")
		s := fmt.Sprintf(a)
		paragraph.Set("innerHTML", s)
		terminal.Call("appendChild", paragraph)
	}

	pkg.Set("value", "")
	condition.Set("value", "")
	return nil
}

func createElement(elementName string) js.Value {
	return document.Call("createElement", elementName)
}

