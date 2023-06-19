package calc

func Factorial(number int) int {
	if number == 0 {
		return 1
	}

	f := 1
	for i := 1; i <= number; i++ {
		f *= i
	}

	return f

}
