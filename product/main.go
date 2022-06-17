package main

import (
	"fmt"
	"github.com/Kwok/microdemo/common"
	"github.com/Kwok/microdemo/product/domain/repository"
	service2 "github.com/Kwok/microdemo/product/domain/service"
	"github.com/Kwok/microdemo/product/handler"
	pb "github.com/Kwok/microdemo/product/proto"
	ot "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4/registry"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	serviceName = "go.micro.service.product"
	version     = "latest"
)

func main() {
	//// 配置consul
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	// 配置jaeger连接
	tracer, closer, err := common.NewJaegerTracer(serviceName, "127.0.0.1:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer closer.Close()
	opentracing.SetGlobalTracer(tracer)

	// Create service
	srv := micro.NewService(
		micro.Name(serviceName),
		micro.Version(version),
		micro.Address("127.0.0.1:8004"),
		micro.WrapHandler(ot.NewHandlerWrapper(tracer)),
		micro.Registry(consulRegistry),
	)
	srv.Init()

	initApp(srv)

	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}

func initApp(srv micro.Service) {
	consulCfg, err := common.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
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

	repo := repository.NewProductRepository(db)
	repo.InitTable()

	productService := service2.NewProductDataService(repo)
	productHandler := handler.New(productService)

	// Register handler
	pb.RegisterProductHandler(srv.Server(), productHandler)
}
