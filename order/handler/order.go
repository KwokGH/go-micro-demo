package handler

import (
	"context"
	"github.com/Kwok/microdemo/common"
	"github.com/Kwok/microdemo/order/domain/model"
	"github.com/Kwok/microdemo/order/domain/service"
	pb "github.com/Kwok/microdemo/order/proto/order"
)

type Order struct {
	OrderDataService service.IOrderDataService
}

func (o *Order) GetOrderByID(ctx context.Context, id *pb.OrderID, info *pb.OrderInfo) error {
	order, err := o.OrderDataService.FindOrderByID(id.OrderId)
	if err != nil {
		return err
	}
	if err := common.Swap(order, info); err != nil {
		return err
	}
	return nil
}

func (o *Order) GetAllOrder(ctx context.Context, request *pb.AllOrderRequest, response *pb.AllOrder) error {
	orderAll, err := o.OrderDataService.FindAllOrder()
	if err != nil {
		return err
	}

	for _, v := range orderAll {
		order := &pb.OrderInfo{}
		if err := common.Swap(v, order); err != nil {
			return err
		}
		response.OrderInfo = append(response.OrderInfo, order)
	}
	return nil
}

func (o *Order) CreateOrder(ctx context.Context, request *pb.OrderInfo, response *pb.OrderID) error {
	orderAdd := &model.Order{}
	if err := common.Swap(request, orderAdd); err != nil {
		return err
	}
	orderID, err := o.OrderDataService.AddOrder(orderAdd)
	if err != nil {
		return err
	}
	response.OrderId = orderID
	return nil
}

func (o *Order) DeleteOrderByID(ctx context.Context, request *pb.OrderID, response *pb.Response) error {
	if err := o.OrderDataService.DeleteOrder(request.OrderId); err != nil {
		return err
	}
	response.Msg = "删除成功"
	return nil
}

func (o *Order) UpdateOrderPayStatus(ctx context.Context, request *pb.PayStatus, response *pb.Response) error {
	if err := o.OrderDataService.UpdatePayStatus(request.OrderId, request.PayStatus); err != nil {
		return err
	}
	response.Msg = "支付状态更新成功"
	return nil
}

func (o *Order) UpdateOrderShipStatus(ctx context.Context, request *pb.ShipStatus, response *pb.Response) error {
	if err := o.OrderDataService.UpdateShipStatus(request.OrderId, request.ShipStatus); err != nil {
		return err
	}
	response.Msg = "发货状态更新成功"
	return nil
}

func (o *Order) UpdateOrder(ctx context.Context, request *pb.OrderInfo, response *pb.Response) error {
	order := &model.Order{}
	if err := common.Swap(request, order); err != nil {
		return err
	}
	if err := o.OrderDataService.UpdateOrder(order); err != nil {
		return err
	}
	response.Msg = "订单更新成功"
	return nil
}

// New Return a new handler
func New(orderService service.IOrderDataService) *Order {
	return &Order{
		orderService,
	}
}
