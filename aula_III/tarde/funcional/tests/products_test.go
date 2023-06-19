package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/anwardh/meliProject/cmd/server/handler"
	"github.com/anwardh/meliProject/internal/products"
	"github.com/anwardh/meliProject/pkg/store"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func createServer() *gin.Engine {
	_ = os.Setenv("TOKEN", "123456")
	db := store.Factory(store.FileType, "../products.json")
	repo := products.NewJsonRepository(db)
	service := products.NewService(repo)
	p := handler.NewProduct(service)
	r := gin.Default()

	pr := r.Group("/products")
	pr.POST("/", p.Store())
	pr.GET("/", p.GetAll())
	return r
}

func createRequestTest(method string, url string, body string) (*http.Request, *httptest.ResponseRecorder) {
	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", "123456")

	return req, httptest.NewRecorder()
}

type Response struct {
	Code string             `json:"code"`
	Data []products.Product `json:"data"`
}

func Test_GetProduct_OK(t *testing.T) {
	// criar um servidor e define suas rotas
	r := createServer()
	// criar uma Request do tipo GET e Response para obter o resultado
	request, response := createRequestTest(http.MethodGet, "/products/", "")
	// diz ao servidor que ele pode atender a solicitação
	r.ServeHTTP(response, request)

	// formato da response
	// {
	// 	code
	// 	Data []bytes `[{ name: "Bolo de cenoura" }]`
	// }

	var responseResult Response
	assert.Equal(t, 200, response.Code)

	err := json.Unmarshal(response.Body.Bytes(), &responseResult)

	assert.Nil(t, err)
	assert.True(t, len(responseResult.Data) > 0)
}

func Test_SaveProduct_OK(t *testing.T) {
	// crie o Servidor e defina as Rotas
	r := createServer()
	// crie Request do tipo POST e Response para obter o resultado
	request, response := createRequestTest(http.MethodPost, "/products/", `{
			"name": "Tester","category": "Funcional","count": 10,"price": 99.99
	}`)

	// diga ao servidor que ele pode atender a solicitação
	r.ServeHTTP(response, request)
	assert.Equal(t, 201, response.Code)
}

// controller <-> service <-> repository <-> story

func Test_SaveProduct_InvalidBody(t *testing.T) {
	// crie o Servidor e defina as Rotas
	r := createServer()
	// crie Request do tipo POST e Response para obter o resultado
	request, response := createRequestTest(http.MethodPost, "/products/", `{
			"name": "","category": "Funcional","count": 10,"price": 99.99
	}`)

	// diga ao servidor que ele pode atender a solicitação
	r.ServeHTTP(response, request)
	assert.Equal(t, 400, response.Code)
}
