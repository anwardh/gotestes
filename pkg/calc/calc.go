package calc

func Soma(a, b int) int {
	return a + b
}

func Subtracao(a, b int) int {
	if b < a {
		return 0
	}
	return a - b

}

func Divisao(a, b int) int {
	return a / b
}
