package main

func main() {
	a := 5
	var ponteiro *int = &a
	*ponteiro = 4
	b := &a
	*b = 98
	println(a)
}
