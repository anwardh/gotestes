package products_test

import (
	"testing"

	"github.com/bootcamp-go/desafio-cierre-testing/internal/products"
	mocks "github.com/bootcamp-go/desafio-cierre-testing/tests/products"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetAllBySeller(t *testing.T) {
	t.Run("Deve retornar os produtos ao chamar a repository", func(t *testing.T) {
		expectedProducts := []products.Product{
			{
				ID:          "1",
				SellerID:    "123",
				Description: "descrição qualquer",
				Price:       20.0,
			},
			{
				ID:          "2",
				SellerID:    "123",
				Description: "descrição qualquer 2",
				Price:       40.0,
			},
		}

		service, repository := CreateService(t)
		repository.On("GetAllBySeller", mock.AnythingOfType("string")).Return(expectedProducts, nil)

		products, err := service.GetAllBySeller("123")

		assert.True(t, len(products) == 2)
		assert.NoError(t, err)
	})

	t.Run("Deve retornar o erro caso ocorra um erro na chamada da repository", func(t *testing.T) {
		emptyProducts := make([]products.Product, 0)
		expectedMessage := "error to connect database"

		service, repository := CreateService(t)
		repository.On("GetAllBySeller", mock.AnythingOfType("string")).Return(emptyProducts, products.ErrConnectDatabase)

		_, err := service.GetAllBySeller("123")

		assert.Error(t, err)
		assert.Equal(t, expectedMessage, err.Error())
	})

}

func CreateService(t *testing.T) (products.Service, *mocks.ProductRepositoryMock) {
	t.Helper()
	repoMock := new(mocks.ProductRepositoryMock)
	service := products.NewService(repoMock)
	return service, repoMock
}
