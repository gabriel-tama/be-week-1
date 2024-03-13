package models

import (
	"context"

	"github.com/jackc/pgx/v5"
)

type BankAccount struct {
	BankAccountId     uint
	UserId            uint
	BankName          string
	BankAccountName   string
	BankAccountNumber string
}

type BankAccountModel struct {
	db *pgx.Conn
}

func NewBankAccountModel(db *pgx.Conn) *BankAccountModel {
	return &BankAccountModel{db: db}
}

func (bm *BankAccountModel) Create(user_id int, bank_name, account_name, account_number string) (*BankAccount,error) {
	// Store the user in the database
    _, err := bm.db.Exec(context.Background(), "INSERT INTO bankaccounts (user_id, bank_name, account_name,account_number) VALUES ($1, $2, $3, $4)",
        user_id,bank_name,account_name,account_number)
    if err != nil {
        return nil,err
    }
	
		bankAcc := &BankAccount{
        BankName: bank_name,
        BankAccountName:     account_name,
		BankAccountNumber: account_number,
    }

	return bankAcc,nil

}

// func (bm *BankAccountModel) FindByUserID(user_id int) (*BankAccount,error)
