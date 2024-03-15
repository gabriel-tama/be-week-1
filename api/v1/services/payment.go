package services

import (
	"github.com/gabriel-tama/be-week-1/models"
	"github.com/gabriel-tama/be-week-1/types"
)

type PaymentService interface {
	CreatePayment(user_id int, productId int, payment types.PaymentCreate) (bool, error)
}

type paymentServiceImpl struct {
	paymentModel *models.PaymentModel
}

func NewPaymentService(paymentModel *models.PaymentModel) PaymentService {
	return &paymentServiceImpl{paymentModel: paymentModel}
}

func (ps *paymentServiceImpl) CreatePayment(user_id int, productId int, payment types.PaymentCreate) (bool, error) {
	exist, err := ps.paymentModel.Create(user_id, productId, payment)
	if err != nil {
		return false, err
	}
	return exist, nil
}