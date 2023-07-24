package repositories

import (
	"example/models"
	"sync"
)

type ProductRepositoryInterface interface {
	Add(product models.Product) error //Add products to data
}

type ProductRepository struct {
	data  map[string]models.Product //ProductID
	mutex sync.Mutex
}

func NewProductRepository() ProductRepository {
	return ProductRepository{
		data:  make(map[string]models.Product),
		mutex: sync.Mutex{},
	}
}

func (r *ProductRepository) Add(product models.Product) error {
	r.mutex.Lock()
	defer r.mutex.Unlock()
	r.data[product.ID] = product
	return nil
}
