package testutil

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"

	"github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {
	// Aqui estamos descartando o erro
	_ = os.Setenv("TOKEN", "123456")
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	return r
}

func MakeRequest(
	method,
	url,
	body string,
) (
	*http.Request,
	*httptest.ResponseRecorder,
) {

	req := httptest.NewRequest(method, url, bytes.NewBuffer([]byte(body)))
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("TOKEN", "123456")

	return req, httptest.NewRecorder()

}
