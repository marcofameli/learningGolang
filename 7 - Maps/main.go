package main

import "fmt"

func main() {
	salarios := map[string]int{"Marco": 1999, "Roger": 3, "Riggins": 2000}
	delete(salarios, "Marco")
	salarios["Ror"] = 5000

	//sal := make(map[string]int)
	//sall := map[string]int{}
	//sall["Wesley"] = 1000

	for _, salario := range salarios {
		fmt.Printf(" O salario é %d\n", salario)
	}

}
