package calc_test

import (
	"testing"

	"github.com/anwardh/gotestes/pkg/calc"
	"github.com/stretchr/testify/assert"
)

func TestSoma(t *testing.T) {
	resultadoAtual := calc.Soma(5, 5)
	resultadoEsperado := 10

	// if resultadoAtual != resultadoEsperado {
	// 	t.Errorf("failed: o resultado atual %v é diferente do resultado esperado %v", resultadoAtual, resultadoEsperado)
	// }

	assert.Equal(t, resultadoEsperado, resultadoAtual, "o resultado atual precisa ser igual ao resultado esperado")
}

func TestSubtracao(t *testing.T) {
	resultadoAtual := calc.Subtracao(5, 3)
	resultadoEsperado := 0

	// if resultadoAtual != resultadoEsperado {
	// 	t.Errorf("failed: o resultado atual %v é diferente do resultado esperado %v", resultadoAtual, resultadoEsperado)
	// }

	assert.Equal(t, resultadoEsperado, resultadoAtual, "o resultado atual precisa ser igual ao resultado esperado")
}

func TestSubtracao2(t *testing.T) {
	resultadoAtual := calc.Subtracao(5, 5)
	resultadoEsperado := 0

	// if resultadoAtual != resultadoEsperado {
	// 	t.Errorf("failed: o resultado atual %v é diferente do resultado esperado %v", resultadoAtual, resultadoEsperado)
	// }

	assert.Equal(t, resultadoEsperado, resultadoAtual, "o resultado atual precisa ser igual ao resultado esperado")
}

func TestDivisao(t *testing.T) {
	resultadoAtual := calc.Divisao(5, 5)

	assert.NotNil(t, resultadoAtual)
}
