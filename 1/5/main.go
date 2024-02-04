package main

import "fmt"

const a = "hello friend"

type ID int

var (
	b bool    = true
	c int     = 10
	d string  = "hello friend"
	e float64 = 1.2
	f ID      = 1
)

func main() {
	var testeArray [3]int
	testeArray[0] = 10
	testeArray[1] = 20
	testeArray[2] = 30

	for i, v := range testeArray {
		fmt.Printf("o valor de indice %d é valor de %d\n", i, v)
	}
}
