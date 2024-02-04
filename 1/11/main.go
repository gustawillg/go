package main

import "fmt"

type Client struct {
	Nome  string
	Idade int
	Ativo bool
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
