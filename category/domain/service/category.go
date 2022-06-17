package service

import (
	"github.com/Kwok/microdemo/category/domain/model"
	"github.com/Kwok/microdemo/category/domain/repository"
)

type ICategoryDataService interface {
	Create(Category *model.Category) (int64, error)
	Update(Category *model.Category) error
	Delete(id string) error
	GetByCategoryName(name string) (*model.Category, error)
}

func NewCategoryDataService(CategoryRepo repository.ICategoryRepository) ICategoryDataService {
	return &CategoryDataService{
		CategoryRepo: CategoryRepo,
	}
}

type CategoryDataService struct {
	CategoryRepo repository.ICategoryRepository
}

func (u *CategoryDataService) Create(category *model.Category) (int64, error) {
	return u.CategoryRepo.Create(category)
}

func (u *CategoryDataService) Update(category *model.Category) error {
	//TODO implement me
	panic("implement me")
}

func (u *CategoryDataService) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func (u *CategoryDataService) GetByCategoryName(name string) (*model.Category, error) {
	return u.CategoryRepo.GetByName(name)
}
