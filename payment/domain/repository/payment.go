package repository

import (
	"github.com/Kwok/microdemo/payment/domain/model"
	"gorm.io/gorm"
)

type IPaymentRepository interface {
	InitTable() error
	FindPaymentByID(int64) (*model.Payment, error)
	CreatePayment(*model.Payment) (int64, error)
	DeletePaymentByID(int64) error
	UpdatePayment(*model.Payment) error
	FindAll() ([]*model.Payment, error)
}

// NewPaymentRepository 创建paymentRepository
func NewPaymentRepository(db *gorm.DB) IPaymentRepository {
	return &PaymentRepository{mysqlDb: db}
}

type PaymentRepository struct {
	mysqlDb *gorm.DB
}

func (u *PaymentRepository) FindPaymentByID(id int64) (*model.Payment, error) {
	var data *model.Payment
	return data, u.mysqlDb.First(data, id).Error
}

func (u *PaymentRepository) CreatePayment(payment *model.Payment) (int64, error) {
	return payment.ID, u.mysqlDb.Create(payment).Error
}

func (u *PaymentRepository) DeletePaymentByID(paymentID int64) error {
	return u.mysqlDb.Where("id = ?", paymentID).Delete(&model.Payment{}).Error
}

func (u *PaymentRepository) UpdatePayment(payment *model.Payment) error {
	return u.mysqlDb.Model(payment).Updates(payment).Error
}

func (u *PaymentRepository) FindAll() (paymentAll []*model.Payment, err error) {
	return paymentAll, u.mysqlDb.Find(&paymentAll).Error
}

// InitTable 初始化表
func (u *PaymentRepository) InitTable() error {
	return u.mysqlDb.AutoMigrate(&model.Payment{})
}
