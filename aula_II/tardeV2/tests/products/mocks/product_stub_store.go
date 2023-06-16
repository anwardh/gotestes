package tests_mocks

import (
	"encoding/json"

	"github.com/anwardh/meliProject/internal/products"
)

type StubProductsStore struct {
}

func (s *StubProductsStore) Write(data interface{}) error {
	return nil
}
func (s *StubProductsStore) Read(data interface{}) error {
	products := []products.Product{}
	jsonData, _ := json.Marshal(products)
	return json.Unmarshal(jsonData, &data)
}
