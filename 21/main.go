package main

import (
	"fmt"
	"matematica/matematica"
)

func main() {
	s := matematica.Soma(10, 20)
	carro := matematica.Carro{Marca: "Fiat"}
	fmt.Println(carro.Marca)
	fmt.Println("resultado:", s)
	fmt.Println(matematica.A)
}
