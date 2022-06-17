package handler

import (
	"fmt"
	"github.com/Kwok/microdemo/common/log"
	"github.com/Kwok/microdemo/payment/proto/payment"
	"github.com/gin-gonic/gin"
	"github.com/plutov/paypal/v3"
	"net/http"
	"strconv"
)

var (
	sandboxClientID = "paypal提供的clientID"
)

type PaymentHandler struct {
	paymentService payment.PaymentService
}

func (a *PaymentHandler) InitRouter(router *gin.Engine) {
	router.GET("/payment/payPalRefund", a.PayPalRefund)
}

func New(paymentService payment.PaymentService) *PaymentHandler {
	return &PaymentHandler{paymentService: paymentService}
}

func (e *PaymentHandler) PayPalRefund(c *gin.Context) {
	ctx := c.Request.Context()

	paymentID, err := strconv.ParseInt(c.Query("payment_id"), 10, 64)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	refundID := c.Query("refund_id")
	if refundID == "" {
		c.JSON(http.StatusBadRequest, nil)
		return
	}
	money := c.Query("money")

	paymentInfo, err := e.paymentService.FindPaymentByID(ctx, &payment.PaymentID{PaymentId: paymentID})
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	sandBox := paypal.APIBaseSandBox
	payout := paypal.Payout{
		SenderBatchHeader: &paypal.SenderBatchHeader{
			EmailSubject:  fmt.Sprintf("%s %s", refundID, "kwok 提醒你收款！"),
			EmailMessage:  fmt.Sprintf("%s %s", refundID, "你有一个收款信息！"),
			SenderBatchID: refundID,
		},
		Items: []paypal.PayoutItem{
			{
				RecipientType: "EMAIL",
				//RecipientWallet: "",
				Receiver: "sb-a1ddg16133033@personal.example.com",
				Amount: &paypal.AmountPayout{
					//币种
					Currency: "USD",
					Value:    money,
				},
				Note:         refundID,
				SenderItemID: refundID,
			},
		},
	}
	//paymentInfo.PaymentSid = "密钥"
	paypalClient, err := paypal.NewClient(sandboxClientID, paymentInfo.PaymentSid, sandBox)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	_, err = paypalClient.GetAccessToken()
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	paymentResult, err := paypalClient.CreateSinglePayout(payout)
	if err != nil {
		log.Error(err)
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	log.Info(paymentResult)

	c.JSON(http.StatusOK, refundID+" 支付成功！")
}
