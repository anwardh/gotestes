package tests_mocks

import (
	"encoding/json"

	"github.com/anwardh/meliProject/internal/products"
)

type StubProducts struct {
}

func (s *StubProducts) Write(data interface{}) error {
	return nil
}

func (s *StubProducts) Read(data interface{}) error {
	products := []products.Product{
		{
			ID:       1,
			Name:     "CellPhone",
			Category: "Tech",
			Count:    3,
			Price:    250,
		}, {
			ID:       2,
			Name:     "Notebook",
			Category: "Tech",
			Count:    10,
			Price:    1750.5,
		},
	}
	jsonData, _ := json.Marshal(products)
	return json.Unmarshal(jsonData, &data)
}
