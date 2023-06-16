package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetVersion(t *testing.T) {
	myDummyDB := DummyRepository{}
	engine := NewService(myDummyDB)
	expectedVersion := "1.0"

	result := engine.GetVersion()

	assert.Equal(t, expectedVersion, result)
}
