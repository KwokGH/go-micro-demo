package handler

import (
	"context"
	"github.com/Kwok/microdemo/common"
	"github.com/Kwok/microdemo/common/log"
	"github.com/Kwok/microdemo/payment/domain/model"
	"github.com/Kwok/microdemo/payment/domain/service"
	pb "github.com/Kwok/microdemo/payment/proto/payment"
)

type Payment struct {
	PaymentDataService service.IPaymentDataService
}

func (p *Payment) AddPayment(ctx context.Context, request *pb.PaymentInfo, response *pb.PaymentID) error {
	payment := &model.Payment{}
	if err := common.Swap(request, payment); err != nil {
		log.Error(err)
	}
	paymentID, err := p.PaymentDataService.AddPayment(payment)
	if err != nil {
		log.Error(err)
	}
	response.PaymentId = paymentID
	return nil
}

func (p *Payment) UpdatePayment(ctx context.Context, request *pb.PaymentInfo, response *pb.Response) error {
	payment := &model.Payment{}
	if err := common.Swap(request, payment); err != nil {
		log.Error(err)
	}

	response.Msg = "更新成功"

	return p.PaymentDataService.UpdatePayment(payment)
}

func (p *Payment) DeletePaymentByID(ctx context.Context, request *pb.PaymentID, response *pb.Response) error {
	if err := p.PaymentDataService.DeletePayment(request.PaymentId); err != nil {
		return err
	}

	response.Msg = "更新成功"

	return nil
}

func (p *Payment) FindPaymentByID(ctx context.Context, request *pb.PaymentID, response *pb.PaymentInfo) error {
	payment, err := p.PaymentDataService.FindPaymentByID(request.PaymentId)
	if err != nil {
		log.Error(err)
	}
	return common.Swap(payment, response)
}

func (p *Payment) FindAllPayment(ctx context.Context, all *pb.All, response *pb.PaymentAll) error {
	allPayment, err := p.PaymentDataService.FindAllPayment()
	if err != nil {
		log.Error(err)
	}

	for _, v := range allPayment {
		paymentInfo := &pb.PaymentInfo{}
		if err := common.Swap(v, paymentInfo); err != nil {
			log.Error(err)
		}
		response.PaymentInfo = append(response.PaymentInfo, paymentInfo)
	}
	return nil
}

// New Return a new handler
func New(paymentService service.IPaymentDataService) *Payment {
	return &Payment{
		paymentService,
	}
}
