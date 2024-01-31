package main

import (
	"fmt"
)

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5, 8, 9, 865, 6, 87, 748, 6541, 5496))
}

func sum(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}
	return total
}
