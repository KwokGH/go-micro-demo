package service

import (
	"github.com/Kwok/microdemo/product/domain/model"
	"github.com/Kwok/microdemo/product/domain/repository"
)

type IProductDataService interface {
	Create(Product *model.Product) (int64, error)
	Update(Product *model.Product) error
	Delete(id int64) error
	GetByID(name int64) (*model.Product, error)
}

func NewProductDataService(ProductRepo repository.IProductRepository) IProductDataService {
	return &ProductDataService{
		ProductRepo: ProductRepo,
	}
}

type ProductDataService struct {
	ProductRepo repository.IProductRepository
}

func (u *ProductDataService) Create(product *model.Product) (int64, error) {
	return u.ProductRepo.Create(product)
}

func (u *ProductDataService) Update(product *model.Product) error {
	return u.ProductRepo.Update(product)
}

func (u *ProductDataService) Delete(id int64) error {
	return u.ProductRepo.Delete(id)
}

func (u *ProductDataService) GetByID(id int64) (*model.Product, error) {
	return u.ProductRepo.GetByID(id)
}
