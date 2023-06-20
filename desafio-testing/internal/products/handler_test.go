package products_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/bootcamp-go/desafio-cierre-testing/internal/products"
	"github.com/bootcamp-go/desafio-cierre-testing/pkg/testutil"
	mocks "github.com/bootcamp-go/desafio-cierre-testing/tests/products"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

const (
	GetProducts = "/api/v1/products"
)

func TestGetAllProducts(t *testing.T) {
	emptyProducts := make([]products.Product, 0)

	t.Run("Deve retornar status code 400 se o query param seller_id não for passado", func(t *testing.T) {
		server, _ := InitServerWithGetProductsRoute(t)

		request, response := testutil.MakeRequest(http.MethodGet, GetProducts, "")

		// Aqui passamos a request para ser processada pelo servidor,
		// e depois disso a variavel response vai ser preenchida
		server.ServeHTTP(response, request)

		assert.Equal(t, response.Code, http.StatusBadRequest)
	})
	t.Run("Deve retornar status code 404 se o seller não existir", func(t *testing.T) {
		server, mockService := InitServerWithGetProductsRoute(t)

		mockService.On(
			"GetAllBySeller",
			mock.AnythingOfType("string")).Return(emptyProducts, products.ErrSellerNotFound)

		request, response := testutil.MakeRequest(
			http.MethodGet,
			GetProducts,
			"",
		)

		DefineQueryParamSellerId(t, request)

		server.ServeHTTP(response, request)

		assert.Equal(t, response.Code, http.StatusNotFound)
	})
	t.Run("Deve retornar status code 204 se o seller não tiver produtos", func(t *testing.T) {
		server, mockService := InitServerWithGetProductsRoute(t)

		mockService.On(
			"GetAllBySeller",
			mock.AnythingOfType("string")).Return(emptyProducts, nil)

		request, response := testutil.MakeRequest(
			http.MethodGet,
			GetProducts,
			"",
		)

		DefineQueryParamSellerId(t, request)

		server.ServeHTTP(response, request)

		assert.Equal(t, response.Code, http.StatusNoContent)
	})
	t.Run("Deve retornar status code 500, caso o método GetAllBySeller da repository retorne um erro", func(t *testing.T) {
		server, mockService := InitServerWithGetProductsRoute(t)

		mockService.On(
			"GetAllBySeller",
			mock.AnythingOfType("string")).Return(emptyProducts, products.ErrConnectDatabase)

		request, response := testutil.MakeRequest(
			http.MethodGet,
			GetProducts,
			"",
		)

		DefineQueryParamSellerId(t, request)

		server.ServeHTTP(response, request)

		assert.Equal(t, http.StatusInternalServerError, response.Code)
	})
	t.Run("Deve retornar status code 200, com os produtos como resposta", func(t *testing.T) {
		server, mockService := InitServerWithGetProductsRoute(t)

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

		request, response := testutil.MakeRequest(
			http.MethodGet,
			GetProducts,
			"",
		)

		DefineQueryParamSellerId(t, request)

		mockService.On(
			"GetAllBySeller",
			mock.AnythingOfType("string")).Return(expectedProducts, nil)

		server.ServeHTTP(response, request)

		responseResult := make([]products.Product, 0)
		_ = json.Unmarshal(response.Body.Bytes(), &responseResult)

		assert.Equal(t, response.Code, http.StatusOK)
		assert.True(t, len(responseResult) == 2)
	})
}

func InitServerWithGetProductsRoute(t *testing.T) (*gin.Engine, *mocks.ProductServiceMock) {
	t.Helper()
	server := testutil.CreateServer()
	mockService := new(mocks.ProductServiceMock)
	handler := products.NewHandler(mockService)
	server.GET(GetProducts, handler.GetProducts)
	return server, mockService
}

func DefineQueryParamSellerId(t *testing.T, request *http.Request) {
	t.Helper()
	// Estamos pegando a query param da url
	query := request.URL.Query() // /api/v1/products
	query.Add("seller_id", "123")
	// A função encode deixa os query params que foram definidos, no formato que o navegador entenda
	request.URL.RawQuery = query.Encode()
	// /api/v1/products?seller_id=123
}
