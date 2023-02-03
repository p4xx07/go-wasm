package main

import (
	"syscall/js"
)

func main() {
	c := make(chan struct{}, 0)

	println("WASM Go Initialized")
	// register functions
	register()
	<-c
}

func register() {
	js.Global().Set("add", js.FuncOf(add))
}

func add(this js.Value, args []js.Value) any {
	a := args[0].Int()
	b := args[1].Int()
	return a + b
}
