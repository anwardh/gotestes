package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearchByNameFaked(t *testing.T) {
	myFakeRepository := FakeRepository{}
	service := NewService(&myFakeRepository)
	testValues := map[string]string{
		"Nacho": "123456",
		"NicoA": "234567",
	}

	for key := range testValues {
		service.AddEntry(key, testValues[key])
	}

	resultNacho := service.SearchByName("Nacho")
	resultNico := service.SearchByName("NicoA")
	resultPhone := service.SearchByPhone("123456")

	assert.Equal(t, testValues["Nacho"], resultNacho)
	assert.Equal(t, testValues["NicoA"], resultNico)
	assert.Equal(t, "Nacho", resultPhone)
}
