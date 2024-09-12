package main

type MyNumber int

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
		return true // PACOTE CONSTRATINS COMPARA COISAS
	}
	return false
}

func main() {
	m := map[string]int{"Marco": 200, "Marcoo": 200, "Marrcoo": 200}
	m2 := map[string]float64{"Marco": 2.00, "Marcoo": 2.00, "Marrcoo": 2.00}
	m3 := map[string]float64{"Marco": 2.00, "Marcoo": 2.00, "Marrcoo": 2.00}
	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10.0))
}
