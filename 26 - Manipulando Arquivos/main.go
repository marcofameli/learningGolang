package main

import (
	"bufio"
	"fmt"
	"os"
)

// ///////////////////// Criar o arquivo
func main() {
	f, err := os.Create("arquivoteste.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := f.Write([]byte("Arquivos\n"))
	//tamanho, err := f.WriteString("Hello, World!")
	if err != nil {
		panic(err)
	}

	fmt.Printf("Arquivo criado com sucesso%d\n", tamanho)
	f.Close()

	/// Leitura
	arquivo, err := os.ReadFile("arquivoteste.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	// Leitura de pouco em pouco lendo o arquivo
	arquivo2, err := os.Open("arquivoteste.txt")
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 10)
	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}
		fmt.Println(string(buffer[:n]))
	}

	err = os.Remove("arquivoteste.txt")
	if err != nil {
		panic(err)
	}

}
