package main

import "fmt"

// DEFER DEIXA A EXECUÇÃO POR ULTIMO
func main() {
	defer fmt.Println("primeira linha")
	fmt.Println("segunda linha")
	fmt.Println("terceira linha")
}
