package main

import (
	"fmt"
	common2 "github.com/Kwok/microdemo/common"
	"github.com/Kwok/microdemo/order/domain/repository"
	"github.com/Kwok/microdemo/order/domain/service"
	"github.com/Kwok/microdemo/order/handler"
	pb "github.com/Kwok/microdemo/order/proto/order"
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4/registry"
	"gorm.io/gorm"

	ot "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	"github.com/go-micro/plugins/v4/wrapper/monitoring/prometheus"
	ratelimit "github.com/go-micro/plugins/v4/wrapper/ratelimiter/uber"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
	"gorm.io/driver/mysql"
)

var (
	serviceName = "go.micro.service.order"
	version     = "latest"
	QPS         = 100
)

func main() {
	//注册中心
	consulRegistry := consul.NewRegistry()
	consulRegistry.Init(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})

	//链路追踪
	t, io, err := common2.NewJaegerTracer(serviceName, "localhost:6831")
	if err != nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// 5.暴露监控地址
	common2.StartPrometheus(9092)

	// Create service
	srv := micro.NewService(
		micro.Name(serviceName),
		micro.Version(version),
		//暴露的服务地址
		micro.Address("0.0.0.0:8005"),
		//注册中心
		micro.Registry(consulRegistry),
		//链路追踪
		micro.WrapHandler(ot.NewHandlerWrapper(opentracing.GlobalTracer())),
		//添加限流
		micro.WrapHandler(ratelimit.NewHandlerWrapper(QPS)),
		//添加监控
		micro.WrapHandler(prometheus.NewHandlerWrapper()),
	)
	srv.Init()

	// Register handler
	initApp(srv)
	// Run service
	if err := srv.Run(); err != nil {
		log.Fatal(err)
	}
}

func initApp(srv micro.Service) {
	consulCfg, err := common2.GetConsulConfig("127.0.0.1", 8500, "/micro/config")
	if err != nil {
		fmt.Println(err)
		return
	}

	mysqlCfg, err := common2.GetMysqlConfig(consulCfg, "mysql")
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

	repo := repository.NewOrderRepository(db)
	repo.InitTable()

	productService := service.NewOrderDataService(repo)
	productHandler := handler.New(productService)

	// Register handler
	pb.RegisterOrderHandler(srv.Server(), productHandler)
}
