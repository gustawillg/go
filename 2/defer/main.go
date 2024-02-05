package main

import "fmt"

func main() {
	defer fmt.Println("primeira linha")
	fmt.Println("segunda linha")
	fmt.Println("terceira linha")
}
