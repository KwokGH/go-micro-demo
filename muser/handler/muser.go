package handler

import (
	"context"
	"github.com/Kwok/microdemo/muser/domain/model"
	"github.com/Kwok/microdemo/muser/domain/service"
	pb "github.com/Kwok/microdemo/muser/proto/user"
)

type User struct {
	UserService service.IUserDataService
}

func (u *User) Login(ctx context.Context, request *pb.UserLoginRequest, response *pb.UserLoginResponse) error {
	userID, err := u.UserService.CheckPwd(request.UserName, request.Password)
	if err != nil {
		return err
	}

	response.UserId = userID

	return nil
}

func (u *User) Register(ctx context.Context, request *pb.UserRegisterRequest, response *pb.UserRegisterResponse) error {
	user := &model.User{
		Name:     request.GetUserName(),
		Password: request.GetPassword(),
	}

	userID, err := u.UserService.Create(user)
	if err != nil {
		return err
	}

	response.UserId = userID

	return nil
}

func (u *User) Get(ctx context.Context, request *pb.UserInfoRequest, response *pb.UserInfoResponse) error {
	userInfo, err := u.UserService.GetByUserName(request.UserName)
	if err != nil {
		return err
	}

	response.Id = userInfo.ID
	response.UserName = userInfo.Name

	return nil
}

// New Return a new handler
func New(userService service.IUserDataService) *User {
	return &User{
		userService,
	}
}
