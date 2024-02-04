package main

import "fmt"

func main() {
	salarios := map[string]int{"Test": 1000, "madoka": 2000}
	delete(salarios, "Test")
	salarios["megumo"] = 5000

	//	sal := make(map[string]int)
	//	sal1 := map[string]int{}
	//	sal1["madoka"] = 2000

	for nome, salario := range salarios {
		fmt.Printf("o salario de  %s é %d\n", nome, salario)
	}

	for _, salario := range salarios {
		fmt.Printf("o salario de é %d\n", salario)
	}
}
