package main

import "fmt"

func main() {
	// Memoria -> Endereço -> Valor

	// Variavel -> ponteiro que tem um endereco na memoria -> valor
	a := 20
	ponteiro := &a
	b := &a
	*b = 30
	fmt.Println(a)
	fmt.Println(*b)
	fmt.Println(ponteiro)
	// o asterisco mostra o valor que ta , sem o asterisco mostra o valor memoria ex: 0000xc93493
	//o & imprime o valor da memoria tambem, sem ele mostra a variavel normal

}
