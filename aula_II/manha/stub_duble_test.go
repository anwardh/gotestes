package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindByName(t *testing.T) {
	myStubRepository := StubRepository{}
	engine := NewService(myStubRepository)
	expectedPhone := "12345678"

	result := engine.FindByName("Pepe")

	assert.Equal(t, expectedPhone, result)
}
