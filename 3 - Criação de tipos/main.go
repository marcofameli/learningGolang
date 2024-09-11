package main

import (
	"fmt"
)

const a = "Hello, World"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "Marco"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	fmt.Printf("o Tipo de E é %T", f) // O T imprime o tipo da variavel
}
