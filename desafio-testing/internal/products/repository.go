package products

import "errors"

type Repository interface {
	GetAllBySeller(sellerID string) ([]Product, error)
}

var (
	ErrConnectDatabase = errors.New("error to connect database")
)

type repository struct{}

func NewRepository() Repository {
	return &repository{}
}

func (r *repository) GetAllBySeller(sellerID string) ([]Product, error) {
	var prodList []Product
	prodList = append(prodList, Product{
		ID:          "mock",
		SellerID:    "FEX112AC",
		Description: "generic product",
		Price:       123.55,
	})
	return prodList, nil
}
