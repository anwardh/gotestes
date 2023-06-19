package main

func TypeNumber(a int) string {
	switch {
	case a < 0:
		return "negativo"
	case a > 0:
		return "positivo"
	default:
		return "zero"
	}
}
