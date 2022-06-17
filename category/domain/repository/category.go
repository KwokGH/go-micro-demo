package repository

import (
	"fmt"
	"github.com/Kwok/microdemo/category/domain/model"
	"gorm.io/gorm"
)

type ICategoryRepository interface {
	InitTable() error
	Create(category *model.Category) (int64, error)
	Update(category *model.Category) error
	Delete(id string) error
	GetByName(name string) (*model.Category, error)
}

func NewCategoryRepository(db *gorm.DB) ICategoryRepository {
	return &CategoryRepository{
		mysqlDB: db,
	}
}

type CategoryRepository struct {
	mysqlDB *gorm.DB
}

func (r *CategoryRepository) InitTable() error {
	return r.mysqlDB.AutoMigrate(&model.Category{})
}

func (r *CategoryRepository) Create(category *model.Category) (int64, error) {
	return category.ID, r.mysqlDB.Create(category).Error
}

func (r *CategoryRepository) Update(category *model.Category) error {
	return r.mysqlDB.Updates(category).Error
}

func (r *CategoryRepository) Delete(id string) error {
	return r.mysqlDB.Delete(&model.Category{}, id).Error
}

func (r *CategoryRepository) FuzzyGetByName(name string) ([]*model.Category, error) {
	var categorys []*model.Category
	err := r.mysqlDB.Where("name like ?", "%"+name+"%").Find(&categorys).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return categorys, nil
}

func (r *CategoryRepository) GetByName(name string) (*model.Category, error) {
	var category *model.Category
	err := r.mysqlDB.Where("name = ?", name).First(&category).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return category, nil
}
