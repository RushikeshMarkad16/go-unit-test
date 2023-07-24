package services

import (
	"errors"
	"example/models"
	"example/repositories"
)

type ProductServiceInterface interface {
	Insert(productID, productName string, price int, stock float64)
}

type ProductService struct {
	repo repositories.ProductRepositoryInterface
}

func NewProductService(repo repositories.ProductRepositoryInterface) ProductService {
	return ProductService{
		repo: repo,
	}
}

func (s ProductService) Insert(productID, productName string, price int, stock float64) error {
	if len(productID) == 0 {
		errors.New("Product ID cannot be null")
	}

	err := s.repo.Add(models.Product{
		ID:    productID,
		Name:  productName,
		Price: float64(price),
		Stock: int(stock),
	})

	if err != nil {
		return err
	}

	return nil
}
