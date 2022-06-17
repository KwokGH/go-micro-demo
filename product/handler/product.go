package handler

import (
	"context"
	pb "github.com/Kwok/microdemo/product/proto"
	log "go-micro.dev/v4/logger"
)

import (
	"github.com/Kwok/microdemo/product/common"
	"github.com/Kwok/microdemo/product/domain/model"
	"github.com/Kwok/microdemo/product/domain/service"
)

type Product struct {
	ProductService service.IProductDataService
}

func (p *Product) AddProduct(ctx context.Context, request *pb.AddProductRequest, response *pb.AddProductResponse) error {
	product := new(model.Product)

	if err := common.Swap(request, product); err != nil {
		log.Info("222", err)
		return err
	}

	id, err := p.ProductService.Create(product)
	if err != nil {
		return err
	}

	response.ProductId = id

	return nil
}

func (p *Product) FindProductByID(ctx context.Context, id *pb.RequestID, info *pb.ProductInfo) error {
	product, err := p.ProductService.GetByID(id.ProductId)
	if err != nil {
		return err
	}

	if err := common.Swap(info, product); err != nil {
		return err
	}

	return nil
}

func (p *Product) UpdateProduct(ctx context.Context, info *pb.ProductInfo, response *pb.Response) error {
	response.ProductId = 100
	return nil
}

func (p *Product) DeleteProductByID(ctx context.Context, id *pb.RequestID, response *pb.Response) error {
	response.ProductId = 100
	return nil
}

// New Return a new handler
func New(productService service.IProductDataService) *Product {
	return &Product{
		productService,
	}
}
