package main

func main() {
	//	for i := 0; i < 10; i++ {
	//		println(i)
	//	}

	numeros := []string{"um", "dois", "tres"}

	for k, v := range numeros {
		println(k, v)
	}
}
