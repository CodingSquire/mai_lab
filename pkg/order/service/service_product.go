package repository

import (
	"go_services_lab/models"
	"go_services_lab/pkg/order/repository"
)

type ProductService struct {
	rep repository.Product
}

func NewProductService(rep repository.Product) *ProductService {
	return &ProductService{rep: rep}
}

func (s *ProductService) Create(product models.Product) (int, error) {
	return s.rep.Create(product)
}

func (s *ProductService) GetAll() ([]models.Product, error) {
	return s.rep.GetAll()
}

func (s *ProductService) LastOne() (models.Product, error) {
	return s.rep.LastOne()
}
