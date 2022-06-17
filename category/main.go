package main

import (
	"fmt"
	"github.com/Kwok/microdemo/category/domain/repository"
	service2 "github.com/Kwok/microdemo/category/domain/service"
	"github.com/Kwok/microdemo/category/handler"
	pb "github.com/Kwok/microdemo/category/proto/category"
	"github.com/Kwok/microdemo/common"
	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	// create a new service
	srv := micro.NewService(
		micro.Name("go.micro.service.category"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8003"),
		micro.Registry(consulRegistry),
	)

	srv.Init()

	initApp(srv)

	// Run service
	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}
}

func initApp(srv micro.Service) {
	consulCfg, err := common.GetConsulConfig("", 8500, "/micro/config")
	if err != nil {
		fmt.Println(err)
		return
	}

	mysqlCfg, err := common.GetMysqlConfig(consulCfg, "mysql")
	if err != nil {
		fmt.Println(err)
		return
	}

	//"root:Passw0rd@tcp(127.0.0.1:3306)/microdemo?charset=utf8mb4&parseTime=True&loc=Local",
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		mysqlCfg.User,
		mysqlCfg.Password,
		mysqlCfg.Host,
		mysqlCfg.Port,
		mysqlCfg.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		return
	}

	repo := repository.NewCategoryRepository(db)
	repo.InitTable()

	categoryService := service2.NewCategoryDataService(repo)
	categoryHandler := handler.New(categoryService)

	// Register handler
	pb.RegisterCategoryHandler(srv.Server(), categoryHandler)
}
