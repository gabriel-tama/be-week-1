package models

import (
	"context"

	"github.com/gabriel-tama/be-week-1/types"
	"github.com/jackc/pgx/v5"
)

type Payment struct {
	AccountID    int
	ProductID    int
	PaymentProof string
	Quantity     int
}

type PaymentModel struct {
	db *pgx.Conn
}

func NewPaymentModel(db *pgx.Conn) *PaymentModel {
	return &PaymentModel{db: db}
}


func (pm *PaymentModel) Create(user_id int, product_id int, payment types.PaymentCreate)(bool,error){
	tx,err := pm.db.Begin(context.Background())

	if err!=nil{
		return false, nil
	}
	defer tx.Rollback(context.Background())

	result,err := tx.Exec(context.Background(), "UPDATE product SET stock= stock - $1 WHERE (id=$2 AND ispurchaseable=true)",
						payment.Quantity,product_id)
	
	if err!=nil{
		return false,err
	}

	rowsAffected := result.RowsAffected()
    if rowsAffected == 0 {
        return false,nil
    }

	result ,err = tx.Exec(context.Background(), 
	"INSERT INTO payment (account_id,product_id,payment_proof,quantity) VALUES ($1,$2,$3,$4) WHERE EXISTS (SELECT 1 FROM bankaccounts AS b WHERE b.id = account_id AND b.is_deleted = FALSE",
	payment.BankAccountId,product_id,payment.PaymentProofImageUrl,payment.Quantity)

	if err != nil {
		return false,err
    }

	rowsAffected = result.RowsAffected()
    if rowsAffected == 0 {
        return false,nil
    }

	
	err = tx.Commit(context.Background())

	if err!=nil{
		return false, err
	}

	return true, nil
}