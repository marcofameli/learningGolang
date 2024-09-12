package main

import "fmt"

type x interface{}

func main() {
	var x interface{} = 10
	var y interface{} = "Hello, World"

	showType(x)
	showType(y)
}
func showType(t interface{}) {
	fmt.Printf("o tipo da variavel é %T e o valor é %v\n", t, t)
}
