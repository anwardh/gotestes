package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchByNameMocked(t *testing.T) {
	myMockRepository := MockRepository{}
	service := NewService(&myMockRepository)
	expectedPhone := "12345678"

	result := service.SearchByName("Nacho")

	assert.Equal(t, expectedPhone, result)
	assert.True(t, myMockRepository.SearchByNameWasCalled)
}

func TestSearchByNameMockedNotCalled(t *testing.T) {
	myMockRepository := MockRepository{}
	service := NewService(&myMockRepository)
	expectedPhone := ""

	// Aqui a função da repository não é invocada
	// por conta que o nome passado é menor que 5 caracteres
	result := service.SearchByName("Nac")

	assert.Equal(t, expectedPhone, result)
	assert.False(t, myMockRepository.SearchByNameWasCalled)
}
