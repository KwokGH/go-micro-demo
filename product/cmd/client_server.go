package main

import (
	"context"
	"fmt"
	"github.com/Kwok/microdemo/common"
	ot "github.com/asim/go-micro/plugins/wrapper/trace/opentracing/v4"
	"github.com/go-micro/plugins/v4/registry/consul"
	"github.com/opentracing/opentracing-go"
	"go-micro.dev/v4"
	"go-micro.dev/v4/registry"
	"log"

	pb "github.com/Kwok/microdemo/product/proto"
)

func main() {
	//注册中心
	consulRegistry := consul.NewRegistry(func(options *registry.Options) {
		options.Addrs = []string{
			"127.0.0.1:8500",
		}
	})
	//链路追踪
	t, io, err := common.NewJaegerTracer("go.micro.service.product.client", "localhost:6831")
	if err != nil {
		log.Fatal(err)
	}
	defer io.Close()
	opentracing.SetGlobalTracer(t)

	service := micro.NewService(
		micro.Name("go.micro.service.product.client"),
		micro.Version("latest"),
		micro.Address("127.0.0.1:8085"),
		//micro.Client(client.NewClient()),
		//添加注册中心
		micro.Registry(consulRegistry),
		//绑定链路追踪 服务端绑定handle 客户端绑定client
		micro.WrapClient(ot.NewClientWrapper(t)),
	)

	// 访问服务器端服务
	productService := pb.NewProductService("go.micro.service.product", service.Client())

	productAdd := &pb.AddProductRequest{
		Name:  "好吃的",
		Sku:   "食品",
		Price: 10,
		Desc:  "product",
		//ProductCategoryId:  1,
		Image: []*pb.ProductImage{
			{
				Name: "image",
				Code: "img_01",
				Url:  "img_url",
			},
			{
				Name: "image",
				Code: "img_01",
				Url:  "img_url",
			},
		},
		Size: []*pb.ProductSize{
			{
				Name: "product-size",
				Code: "product-size-code",
			},
		},
		Seo: &pb.ProductSeo{
			Title:    "seo-title",
			Keywords: "seo-keywords",
			Desc:     "seo-desc",
			Code:     "seo-code",
		},
	}
	response, err := productService.AddProduct(context.TODO(), productAdd)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(response)
}
