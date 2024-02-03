package main

type Mynumber int

type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"test": 1000, "test2": 2000, "test3": 3000, "test4": 4000}
	m2 := map[string]float64{"test": 100.20, "test2": 200.9, "test3": 300.5, "test4": 400.2}
	m3 := map[string]Mynumber{"test": 1000, "test2": 2000, "test3": 3000, "test4": 4000}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10.20))
}
