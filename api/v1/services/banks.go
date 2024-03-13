package services

import "github.com/gabriel-tama/be-week-1/models"

type BankService interface {
	CheckInBank() error
}

type bankServiceImpl struct {
	bankAccountModel *models.BankAccountModel
}

func NewBankService(bankAccountModel *models.BankAccountModel) BankService {
	return &bankServiceImpl{bankAccountModel: bankAccountModel}
}


func (bs *bankServiceImpl) CheckInBank()error{
	return nil
}