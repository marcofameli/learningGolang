package main

import "fmt"

type Endereço struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Pessoa interface { // INTERFACE É APENAS METODOS
	Desativar()
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereço
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {

}

func (c Cliente) Desativar() { // METODOS
	c.Ativo = false
	fmt.Printf("o cliente %s foi desativado ", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	w := Cliente{
		Nome:  "Wesley",
		Idade: 21,
		Ativo: false,
	}

	minhaEmpresa := Empresa{}
	Desativacao(w)
	Desativacao(minhaEmpresa)

}
