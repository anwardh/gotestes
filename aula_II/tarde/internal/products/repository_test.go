package products

import (
	"encoding/json"
	"testing"

	"github.com/anwardh/meliProject/pkg/store"
	"github.com/go-playground/assert/v2"
)

func TestGetAll(t *testing.T) {
	// Initializing input/output
	input := []Product{
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
	dataJson, _ := json.Marshal(input)
	dbMock := store.Mock{
		Data: dataJson,
	}
	storeStub := store.FileStore{
		FileName: "",
		Mock:     &dbMock,
	}

	// poderiamos passar um stub que implementa
	// o método GetAll e que já retorna os dados
	myRepo := NewRepository(&storeStub)
	// Teste da Execução
	resp, _ := myRepo.GetAll()
	// Validação
	assert.Equal(t, input, resp)
}
