package products

import (
	"errors"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service,
	}
}

func (h *Handler) GetProducts(ctx *gin.Context) {
	// /api/v1/products?seller_id=123

	// Quando eu tiver meu mac, para usar o f2, tem que clicar com a tecla FN junto
	sellerID := ctx.Query("seller_id")
	if sellerID == "" {
		ctx.JSON(400, gin.H{"error": "seller_id query param is required"})
		return
	}

	products, err := h.service.GetAllBySeller(sellerID)

	if errors.Is(err, ErrSellerNotFound) {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	if err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	if len(products) == 0 {
		ctx.JSON(204, products)
		return
	}

	ctx.JSON(200, products)

}
