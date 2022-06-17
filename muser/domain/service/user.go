package service

import (
	"github.com/Kwok/microdemo/muser/domain/model"
	"github.com/Kwok/microdemo/muser/domain/repository"
	"github.com/Kwok/microdemo/muser/utils"
)

type IUserDataService interface {
	Create(user *model.User) (int64, error)
	Update(user *model.User) error
	Delete(id string) error
	GetByUserName(name string) (*model.User, error)
	CheckPwd(userName string, pwd string) (int64, error)
}

func NewUserDataService(userRepo repository.IUserRepository) IUserDataService {
	return &UserDataService{
		UserRepo: userRepo,
	}
}

type UserDataService struct {
	UserRepo repository.IUserRepository
}

func (u *UserDataService) Create(user *model.User) (int64, error) {
	pwd, err := utils.GenerateFromPassword(user.Password)
	if err != nil {
		return 0, err
	}

	user.Password = pwd

	return u.UserRepo.Create(user)
}

func (u *UserDataService) Update(user *model.User) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserDataService) Delete(id string) error {
	//TODO implement me
	panic("implement me")
}

func (u *UserDataService) GetByUserName(name string) (*model.User, error) {
	return u.UserRepo.GetByName(name)
}

func (u *UserDataService) CheckPwd(userName string, pwd string) (int64, error) {
	userInfo, err := u.GetByUserName(userName)
	if err != nil {
		return 0, err
	}

	ok, err := utils.ComparePassword(userInfo.Password, pwd)
	if !ok || err != nil {
		return 0, err
	}

	return userInfo.ID, nil
}
