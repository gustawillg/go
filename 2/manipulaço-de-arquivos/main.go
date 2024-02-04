package main

import (
	"bufio"
	"fmt"
	"os"
)

// criando arquivo
func main() {
	f, err := os.Create("arquivo.txt")
	if err != nil {
		panic(err)
	}

	tamanho, err := f.Write([]byte("escrevendo"))
	// tamanho, err := f.WriteString("hello friend")
	if err != nil {
		panic(err)
	}
	fmt.Printf("arquivo criado, tamanho: %d bytes\n", tamanho)
	f.Close()

	// leitura
	arquivo, err := os.ReadFile("arquivo.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(arquivo))

	//leitura de pouco em pouco abrindo o arquivo
	arquivo2, err := os.Open("arquivo.txt")
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

	//remover arquivo

	err = os.Remove("arquivo.txt")
	if err != nil {
		panic(err)
	}
}
