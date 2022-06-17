package handler

import (
	"context"
	"github.com/Kwok/microdemo/category/common"
	"github.com/Kwok/microdemo/category/domain/model"
	"github.com/Kwok/microdemo/category/domain/service"
	pb "github.com/Kwok/microdemo/category/proto/category"
)

type Category struct {
	CategoryService service.ICategoryDataService
}

func (c *Category) CreateCategory(ctx context.Context, request *pb.CreateCategoryRequest, response *pb.CreateCategoryResponse) error {
	category := new(model.Category)
	if err := common.Swap(request, category); err != nil {
		return err
	}

	id, err := c.CategoryService.Create(category)
	if err != nil {
		return err
	}

	response.Id = id

	return nil
}

func (c *Category) GetCategoryByName(ctx context.Context, request *pb.GetCategoryByNameRequest, response *pb.GetCategoryByNameResponse) error {
	category, err := c.CategoryService.GetByCategoryName(request.Name)
	if err != nil {
		return err
	}

	return common.Swap(category, response)
}

// New Return a new handler
func New(categoryService service.ICategoryDataService) *Category {
	return &Category{
		categoryService,
	}
}
