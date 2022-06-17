package repository

import (
	"github.com/Kwok/microdemo/product/domain/model"
	"gorm.io/gorm"
)

type IProductRepository interface {
	InitTable() error
	Create(product *model.Product) (int64, error)
	Update(product *model.Product) error
	Delete(id int64) error
	GetByID(id int64) (*model.Product, error)
}

func NewProductRepository(db *gorm.DB) IProductRepository {
	return &ProductRepository{
		mysqlDB: db,
	}
}

type ProductRepository struct {
	mysqlDB *gorm.DB
}

func (r *ProductRepository) GetByID(id int64) (*model.Product, error) {
	var product = new(model.Product)
	return product, r.mysqlDB.
		Preload("product_image").
		Preload("product_size").
		Preload("product_seo").
		First(product, id).
		Error
}

func (r *ProductRepository) InitTable() error {
	return r.mysqlDB.AutoMigrate(&model.Product{}, &model.ProductImage{}, &model.ProductSeo{}, &model.ProductSize{})
}

func (r *ProductRepository) Create(product *model.Product) (int64, error) {
	return product.ID, r.mysqlDB.Create(product).Error
}

func (r *ProductRepository) Update(product *model.Product) error {
	return r.mysqlDB.Updates(product).Error
}

func (r *ProductRepository) Delete(id int64) error {
	return r.mysqlDB.Transaction(func(tx *gorm.DB) error {
		if err := tx.Where("id = ?", id).Delete(&model.Product{}).Error; err != nil {
			return err
		}
		if err := tx.Where("product_id = ?", id).Delete(&model.ProductImage{}).Error; err != nil {
			return err
		}
		if err := tx.Where("product_id = ?", id).Delete(&model.ProductSeo{}).Error; err != nil {
			return err
		}
		if err := tx.Where("product_id = ?", id).Delete(&model.ProductSize{}).Error; err != nil {
			return err
		}

		return nil
	})
}
