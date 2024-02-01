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

type Pessoa interface {
	Desativar()
}

type Empresa struct {
	Nome string
}

func (e Empresa) Desativar() {
}

func (c Cliente) Desativar() {
	c.Ativo = false
	fmt.Printf("o cliente %s foi desativado", c.Nome)
}

func Desativacao(pessoa Pessoa) {
	pessoa.Desativar()
}

func main() {
	teste := Cliente{
		Nome:  "teste",
		Idade: 5,
		Ativo: true,
	}
	minhaEmpresa := Empresa{}

	Desativacao(minhaEmpresa)
}
