package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

type Client struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	teste := Client{
		Nome:  "teste",
		Idade: 5,
		Ativo: true,
	}
	teste.Ativo = false

	fmt.Println(teste.Nome)
}
