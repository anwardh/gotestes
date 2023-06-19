package calc_test

import (
	"testing"

	calc "github.com/lucasti79/gotdd"
)

func TestFactorial(t *testing.T) {
	tests := []struct {
		arg  int
		want int
	}{{0, 1}, {5, 120}}

	for i, d := range tests {
		result := calc.Factorial(d.arg)
		if result != d.want {
			t.Errorf("Test[%d]: factorial(%d) returned %d, want %d",
				i, d.arg, result, d.want)
		}
	}
}
