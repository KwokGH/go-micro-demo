package repository

import (
	"fmt"
	"github.com/Kwok/microdemo/muser/domain/model"
	"gorm.io/gorm"
)

type IUserRepository interface {
	InitTable() error
	Create(user *model.User) (int64, error)
	Update(user *model.User) error
	Delete(id string) error
	GetByName(name string) (*model.User, error)
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		mysqlDB: db,
	}
}

type UserRepository struct {
	mysqlDB *gorm.DB
}

func (u *UserRepository) InitTable() error {
	return u.mysqlDB.AutoMigrate(&model.User{})
}

func (u *UserRepository) Create(user *model.User) (int64, error) {
	return user.ID, u.mysqlDB.Create(user).Error
}

func (u *UserRepository) Update(user *model.User) error {
	return u.mysqlDB.Updates(user).Error
}

func (u *UserRepository) Delete(id string) error {
	return u.mysqlDB.Delete(&model.User{}, id).Error
}

func (u *UserRepository) FuzzyGetByName(name string) ([]*model.User, error) {
	var users []*model.User
	err := u.mysqlDB.Where("name like ?", "%"+name+"%").Find(&users).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return users, nil
}

func (u *UserRepository) GetByName(name string) (*model.User, error) {
	var user *model.User
	err := u.mysqlDB.Where("name = ?", name).First(&user).Error
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return user, nil
}
