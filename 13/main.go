package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

type Endereco struct {
	Logradouro string
	Numero     int
	Cidade     string
	Estado     string
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("o cliente %s foi desativado", c.Nome)
}

func main() {
	teste := Cliente{
		Nome:  "teste",
		Idade: 5,
		Ativo: true,
	}
	teste.Ativo = false
	teste.Desativar()

	fmt.Println(teste.Nome)
}
