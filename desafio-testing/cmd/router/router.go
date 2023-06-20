package router

import (
	"github.com/bootcamp-go/desafio-cierre-testing/internal/products"
	"github.com/gin-gonic/gin"
)

func MapRoutes(r *gin.Engine) {
	rg := r.Group("/api/v1")
	{
		buildProductsRoutes(rg)
	}
}

func buildProductsRoutes(r *gin.RouterGroup) {
	repo := products.NewRepository()
	service := products.NewService(repo)
	handler := products.NewHandler(service)

	prodRoute := r.Group("/products")
	{
		prodRoute.GET("", handler.GetProducts)
	}

}
