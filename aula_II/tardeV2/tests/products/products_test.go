package tests

import (
	"testing"

	"github.com/anwardh/meliProject/internal/products"
	tests_mocks "github.com/anwardh/meliProject/tests/products/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetAll(t *testing.T) {
	expectedResult := []products.Product{
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

	MyStub := tests_mocks.StubProducts{}
	MyRepoMock := products.NewRepository(&MyStub)
	MyService := products.NewService(MyRepoMock)
	result, err := MyService.GetAll()

	assert.Nil(t, err)
	assert.NoError(t, err)
	assert.Equal(t, result, expectedResult)
}

func TestGetAllError(t *testing.T) {
	MyStub := tests_mocks.StubProductsError{}
	MyRepoMock := products.NewRepository(&MyStub)
	MyService := products.NewService(MyRepoMock)
	result, err := MyService.GetAll()

	expectedError := "erro ao ler os dados"

	assert.Nil(t, result)
	assert.Error(t, err)
	assert.Equal(t, err.Error(), expectedError)
}

func TestStore(t *testing.T) {
	testProduct := products.Product{
		Name:     "Batata",
		Category: "Vegetal",
		Count:    50,
		Price:    5,
	}

	// estamos testando o método Store da service, a integração dele com a repository
	MyStub := tests_mocks.StubProductsStore{}
	MyRepoMock := products.NewRepository(&MyStub)
	MyService := products.NewService(MyRepoMock)
	result, err := MyService.Store(
		testProduct.Name,
		testProduct.Category,
		testProduct.Count,
		testProduct.Price,
	)

	assert.Nil(t, err)
	assert.Equal(t, result.ID, 1)
}

func TestStoreError(t *testing.T) {
	testProduct := products.Product{}
	expectedErrorMessage := "JSON unexpected character"

	// estamos testando o método Store da service, a integração dele com a repository
	MyStub := tests_mocks.StubProductsStoreError{}
	MyRepoMock := products.NewRepository(&MyStub)
	MyService := products.NewService(MyRepoMock)
	_, err := MyService.Store(
		testProduct.Name,
		testProduct.Category,
		testProduct.Count,
		testProduct.Price,
	)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), expectedErrorMessage)
}
