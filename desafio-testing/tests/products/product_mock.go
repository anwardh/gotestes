package mocks

import (
	"github.com/bootcamp-go/desafio-cierre-testing/internal/products"
	"github.com/stretchr/testify/mock"
)

type ProductServiceMock struct {
	mock.Mock
}

func (m *ProductServiceMock) GetAllBySeller(
	sellerID string,
) (
	[]products.Product,
	error,
) {
	args := m.Called(sellerID)
	return args.Get(0).([]products.Product), args.Error(1)
}

type ProductRepositoryMock struct {
	mock.Mock
}

func (m *ProductRepositoryMock) GetAllBySeller(
	sellerID string,
) (
	[]products.Product,
	error,
) {
	args := m.Called(sellerID)
	return args.Get(0).([]products.Product), args.Error(1)
}
