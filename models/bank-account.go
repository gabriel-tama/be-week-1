package models

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type BankAccount struct {
	BankAccountId     string
	// UserId            uint
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

func (bm *BankAccountModel) FindByUserId(user_id int) ([]BankAccount,error){
	var banks []BankAccount

	rows,err := bm.db.Query(context.Background(), "SELECT account_id,bank_name,account_name,account_number FROM bankaccounts WHERE user_id= $1",user_id)
	if err != nil {
		log.Fatal(err)
        return nil, err
    }
    defer rows.Close()


    for rows.Next() {
        var bank BankAccount
        err := rows.Scan(&bank.BankAccountId,&bank.BankName, &bank.BankAccountName, &bank.BankAccountNumber)
        if err != nil {
            return nil, err
        }
        banks = append(banks, bank)
    }
    if err := rows.Err(); err != nil {
        return nil, err
    }
	return banks,nil
}

func (bm *BankAccountModel) Delete(acc_id int,user_id int)(bool,error){
	result, err := bm.db.Exec(context.Background(), "DELETE FROM bankaccounts WHERE user_id=$1 AND account_id=$2",user_id,acc_id)
    if err != nil {
        return false, err
    }

	rowsAffected := result.RowsAffected()
    if rowsAffected == 0 {
        return false,nil
    }
	return true,nil
}


func (bm *BankAccountModel) Update(acc_id int, user_id int, bankName string, bankAccountName string, bankAccountNumber string)(bool,error){
	result,err := bm.db.Exec(context.Background(),"UPDATE bankaccounts SET bank_name = $1, account_name=$2, account_number=$3 WHERE user_id=$4 AND account_id=$5",bankName,bankAccountName,bankAccountNumber,user_id,acc_id)
	if err != nil {
        return false, err
    }

	rowsAffected := result.RowsAffected()
    if rowsAffected == 0 {
        return false,nil
    }
	return true,nil
}