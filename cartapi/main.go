package main

import (
	"context"
	"fmt"
	"github.com/Kwok/microdemo/cart/proto/cart"
	common2 "github.com/Kwok/microdemo/common"
	"github.com/afex/hystrix-go/hystrix"
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4/client"
	"go-micro.dev/v4/server"
	"net"
	"net/http"

	httpServer "github.com/asim/go-micro/plugins/server/http/v4"
	"github.com/asim/go-micro/plugins/wrapper/select/roundrobin/v4"
	ot "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	"github.com/gin-gonic/gin"
	"go-micro.dev/v4"
	log "go-micro.dev/v4/logger"
)

var (
	serviceName = "http.api.cart"
	version     = "latest"
)

func main() {
	// 配置http服务
	srv := httpServer.NewServer(
		server.Name(serviceName),
		server.Address(":8080"),
	)

	// 使用gin作为http服务
	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())

	// consul服务注册
	consulRegistry := consul.NewRegistry()

	//链路追踪
	t, io, err := common2.NewJaegerTracer(serviceName, "localhost:6831")
	if err != nil {
		log.Error(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	// 熔断器
	// 配置
	hystrix.ConfigureCommand("go.micro.service.cart.Cart.GetAll", hystrix.CommandConfig{MaxConcurrentRequests: 5})
	hystrixHandler := hystrix.NewStreamHandler()
	hystrixHandler.Start()
	go func() {
		err := http.ListenAndServe(net.JoinHostPort("0.0.0.0", "9090"), hystrixHandler)
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
	cartService := cart.NewCartService("go.micro.service.cart", microService.Client())
	demo := newDemo(cartService)
	demo.InitRouter(router)

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

//demo
type demo struct {
	cartService cart.CartService
}

func newDemo(cartService cart.CartService) *demo {
	return &demo{
		cartService: cartService,
	}
}

func (a *demo) InitRouter(router *gin.Engine) {
	router.GET("/demo", a.demo)
}

func (a *demo) demo(c *gin.Context) {
	cartList, err := a.cartService.GetAll(context.TODO(), &cart.CartFindAll{UserId: 123})
	if err != nil {
		c.JSON(500, err)
		return
	}
	c.JSON(200, cartList)
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
