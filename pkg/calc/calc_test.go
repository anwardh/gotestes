package calc_test

import (
	"testing"

	"github.com/anwardh/gotestes/pkg/calc"
)

func TestSoma(t *testing.T) {
	resultadoAtual := calc.Soma(5, 5)
	resultadoEsperado := 10

	if resultadoAtual != resultadoEsperado {
		t.Errorf("failed: o resultado atual %v é diferente do resultado esperado %v", resultadoAtual, resultadoEsperado)
	}
}

func TestSubtracao(t *testing.T) {
	resultadoAtual := calc.Subtracao(5, 5)
	resultadoEsperado := 0

	if resultadoAtual != resultadoEsperado {
		t.Errorf("failed: o resultado atual %v é diferente do resultado esperado %v", resultadoAtual, resultadoEsperado)
	}
}
