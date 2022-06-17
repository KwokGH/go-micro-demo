package main

import (
	"context"
	"fmt"
	common2 "github.com/Kwok/microdemo/common"
	"github.com/Kwok/microdemo/payment/proto/payment"
	"github.com/Kwok/microdemo/paymentapi/handler"
	"github.com/afex/hystrix-go/hystrix"
	httpServer "github.com/asim/go-micro/plugins/server/http/v4"
	"github.com/asim/go-micro/plugins/wrapper/select/roundrobin/v4"
	ot "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	"github.com/gin-gonic/gin"
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
	"go-micro.dev/v4/client"
	log "go-micro.dev/v4/logger"
	"go-micro.dev/v4/registry"
	"go-micro.dev/v4/server"
	"net"
	"net/http"
)

var (
	serviceName = "http.api.payment"
	version     = "latest"
)

func main() {
	// 配置http服务
	srv := httpServer.NewServer(
		server.Name(serviceName),
		server.Address(":8081"),
	)

	// 使用gin作为http服务
	gin.SetMode(gin.DebugMode)
	router := gin.New()
	router.Use(gin.Recovery())

	// consul服务注册
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

	// 熔断器
	// 配置
	//hystrix.ConfigureCommand("go.micro.service.cart.Cart.GetAll", hystrix.CommandConfig{MaxConcurrentRequests: 5})
	hystrixHandler := hystrix.NewStreamHandler()
	hystrixHandler.Start()
	go func() {
		err := http.ListenAndServe(net.JoinHostPort("0.0.0.0", "9091"), hystrixHandler)
		if err != nil {
			log.Error("hystrix server error", err)
		}
	}()

	// Create service
	microService := micro.NewService(
		micro.Server(srv),
		// micro.Registry(registry.NewRegistry()),
		micro.Registry(consulRegistry),
		// 链路追踪
		micro.WrapHandler(ot.NewHandlerWrapper(opentracing.GlobalTracer())),
		// 熔断器
		micro.WrapClient(NewClientHystrixWrapper()),
		//添加负载均衡
		micro.WrapClient(roundrobin.NewClientWrapper()),
	)

	// register router
	paymentService := payment.NewPaymentService("go.micro.service.payment", microService.Client())
	paymentHandler := handler.New(paymentService)
	paymentHandler.InitRouter(router)

	hd := srv.NewHandler(router)
	if err := srv.Handle(hd); err != nil {
		log.Fatal(err)
	}

	srv.Init()

	// Run service
	if err := microService.Run(); err != nil {
		log.Fatal(err)
	}
}

type clientWrapper struct {
	client.Client
}

func (c *clientWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	return hystrix.Do(req.Service()+"."+req.Endpoint(), func() error {
		//run 正常执行
		fmt.Println(req.Service() + "." + req.Endpoint())
		return c.Client.Call(ctx, req, rsp, opts...)
	}, func(err error) error {
		fmt.Println(err)
		return err
	})
}
func NewClientHystrixWrapper() client.Wrapper {
	return func(i client.Client) client.Client {
		return &clientWrapper{i}
	}
}
