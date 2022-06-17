package main

import (
	"fmt"
	"github.com/Kwok/microdemo/muser/domain/repository"
	service2 "github.com/Kwok/microdemo/muser/domain/service"
	"github.com/Kwok/microdemo/muser/handler"
	pb "github.com/Kwok/microdemo/muser/proto/user"
	"go-micro.dev/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// create a new service
	srv := micro.NewService(
		micro.Name("go.micro.service.user"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8002"),
	)

	srv.Init()

	initApp(srv)

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

func initApp(srv micro.Service) {
	dsn := "root:Passw0rd@tcp(127.0.0.1:3306)/microdemo?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	repo := repository.NewUserRepository(db)
	repo.InitTable()

	userService := service2.NewUserDataService(repo)
	userHandler := handler.New(userService)

	// Register handler
	pb.RegisterUserHandler(srv.Server(), userHandler)
}
