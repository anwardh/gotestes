package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchByPhone(t *testing.T) {
	mySpyRepository := SpyRepository{SearchByPhoneWasCalled: false}
	service := NewService(&mySpyRepository)
	phone := "12345678" // em nosso mecanismo, não realizaremos pesquisas se o telefone
	// tem um comprimento inferior a 5

	service.SearchByPhone(phone)

	assert.True(t, mySpyRepository.SearchByPhoneWasCalled)
}

func TestSearchByPhoneNotCalled(t *testing.T) {
	mySpyRepository := SpyRepository{}
	service := NewService(&mySpyRepository)

	service.SearchByPhone("1234") // o telefone tem menos de 5 caracteres

	assert.False(t, mySpyRepository.SearchByPhoneWasCalled)
}
