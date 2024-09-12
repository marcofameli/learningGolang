package main

func soma(a, b *int) int {
	*a = 100
	*b = 100
	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20
	soma(&minhaVar1, &minhaVar2)
	println((minhaVar1))
	println((minhaVar2))
	println(soma(&minhaVar1, &minhaVar2))
}
