package services

import "github.com/gabriel-tama/be-week-1/models"

type BankService interface {
	CheckInBank() error
	CreateBankAccount(userId int, bankName, bankAccountName, bankAccountNumber string) (*models.BankAccount, error)
	GetBankAccount(userId int) ([]*models.BankAccount, error)
	DeleteBankAccount(acc_id, user_id int) (bool, error)
	UpdateBankAccount(acc_id, user_id int, bankName string, bankAccountName string, bankAccountNumber string) (bool, error)
}

type bankServiceImpl struct {
	bankAccountModel *models.BankAccountModel
}

func NewBankService(bankAccountModel *models.BankAccountModel) BankService {
	return &bankServiceImpl{bankAccountModel: bankAccountModel}
}

func (bs *bankServiceImpl) CheckInBank() error {
	return nil // ??
}

func (bs *bankServiceImpl) CreateBankAccount(userId int, bankName, bankAccountName, bankAccountNumber string) (*models.BankAccount, error) {
	bank, err := bs.bankAccountModel.Create(userId, bankName, bankAccountName, bankAccountNumber)
	if err != nil {
		return nil, err
	}
	return bank, nil
}

func (bs *bankServiceImpl) GetBankAccount(userId int) ([]*models.BankAccount, error) {
	bank, err := bs.bankAccountModel.FindByUserId(userId)
	if err != nil {
		return nil, err
	}
	var bankPtrs []*models.BankAccount
	for i := range bank {
		bankPtrs = append(bankPtrs, &bank[i])
	}
	return bankPtrs, nil
}

func (bs *bankServiceImpl) DeleteBankAccount(acc_id, user_id int) (bool, error) {

	exists, err := bs.bankAccountModel.Delete(acc_id, user_id)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (bs *bankServiceImpl) UpdateBankAccount(acc_id, user_id int, bankName string, bankAccountName string, bankAccountNumber string) (bool, error) {

	exists, err := bs.bankAccountModel.Update(acc_id, user_id, bankName, bankAccountName, bankAccountNumber)
	if err != nil {
		return false, err
	}
	return exists, nil
}
