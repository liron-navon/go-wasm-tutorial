package main

import (
	"syscall/js"
	"fmt"
)

func main() {
	done := make(chan struct{}, 0)
	js.Global().Set("__wasm_main", js.NewCallback(register))
	fmt.Println("WASM Go Initialized")
	<- done
}

func execCallback(args []js.Value, value interface{}, e error) {
	var last = args[len(args) - 1]
	if last.Type() == js.TypeFunction {
		last.Invoke(e, value)
	} else {
		fmt.Println("no callback")
	}
}

func register(args []js.Value) {
	var scope = args[0]
	scope.Set("add", js.NewCallback(add))
	execCallback(args, scope, nil)
}

func add(args []js.Value) {
	var a = args[0].Int()
	var b = args[1].Int()
	var sum = a + b
	execCallback(args, sum, nil)
}