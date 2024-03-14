package services

import "github.com/gabriel-tama/be-week-1/models"

type BankService interface {
	CheckInBank() error
	CreateBankAccount(userId int, bankName, bankAccountName, bankAccountNumber string)(*models.BankAccount, error)
	GetBankAccount(userId int)([]*models.BankAccount , error)
	DeleteBankAccount(acc_id, user_id int)(error)
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

func (bs *bankServiceImpl) CreateBankAccount(userId int, bankName, bankAccountName, bankAccountNumber string)(*models.BankAccount,error){
	bank,err := bs.bankAccountModel.Create(userId,bankName,bankAccountName,bankAccountNumber)
	if err!=nil {
		return nil,err
	}
	return bank,nil
}

func (bs *bankServiceImpl) GetBankAccount(userId int)([]*models.BankAccount,error){
	bank,err := bs.bankAccountModel.FindByUserId(userId)
	if err!=nil {
		return nil,err
	}
	var bankPtrs []*models.BankAccount
    for i := range bank {
        bankPtrs = append(bankPtrs, &bank[i])
    }
	return bankPtrs,nil
}

func (bs *bankServiceImpl) DeleteBankAccount(acc_id,user_id int)(error){

	err:=bs.bankAccountModel.Delete(acc_id,user_id)
	if err!=nil{
		return err
	}
	return nil
}

