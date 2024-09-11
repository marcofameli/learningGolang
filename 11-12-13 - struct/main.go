package main

import "fmt"

type Endereço struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereço
}

func (c Cliente) Desativar() { // METODOS
	c.Ativo = false
	fmt.Printf("o cliente %s foi desativado ", c.Nome)
}

func main() {
	w := Cliente{
		Nome:  "Wesley",
		Idade: 21,
		Ativo: false}
	w.Ativo = false
	w.Endereço.Cidade = "SP"
	w.Desativar()

}
