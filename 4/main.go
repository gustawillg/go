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
	fmt.Printf("o tipo de E é %T", f)
}
