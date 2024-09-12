package main

import (
	"Goland/21-22-23-pacotesEmodulos/matematica"
	"fmt"
)

func main() {
	s := matematica.Soma(10, 20)
	fmt.Println("result", s)
} //TUDO QUE COMEÇA COM LETRA MAIUSCULA ESTÁ EXPORTADO
// TUDO COM MINUSCULA É PACOTE LOCAL (SOMENTE DENTRO DO PACOTE NÃO ESTARÁ EXPORTADO)
// GO TIDY (COMANDO) VERIFICA OS PACOTES E AJUDA O CODIGO, É O GERENCIADOR
