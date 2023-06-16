package tests_mocks

import (
	"errors"
)

type StubProductsStoreError struct {
}

func (s *StubProductsStoreError) Write(data interface{}) error {
	return nil
}
func (s *StubProductsStoreError) Read(data interface{}) error {
	return errors.New("JSON unexpected character")
}
