package models

import (
	"context"

	"github.com/gabriel-tama/be-week-1/types"
	"github.com/jackc/pgx/v5"
)

type Product struct {
	ID             int
	Name           string
	Price          float64
	ImageURL       string
	Stock          int
	Condition      string
	IsPurchaseable bool
	Tags           []string
}

type ProductModel struct {
	db *pgx.Conn
}

func NewProductModel(db *pgx.Conn) *ProductModel {
	return &ProductModel{db: db}
}

func (pm *ProductModel) Create(user_id int, product types.ProductCreate)(error){
	tx,err := pm.db.Begin(context.Background())

	if err!=nil{
		return nil
	}
	defer tx.Rollback(context.Background())

	var productId int 

	err = tx.QueryRow(context.Background(), "INSERT INTO product (user_id, name, price, imageUrl,stock,condition,isPurchaseable) VALUES ($1, $2, $3, $4,$5,$6,$7) RETURNING id",
	user_id,product.Name,product.Price,product.ImageURL,product.Stock,product.Condition,product.IsPurchaseable).Scan(&productId)

	if err != nil {
		return err
    }

	stmt:= "INSERT INTO product_tags (product_id, tag) VALUES ($1, $2)"
    // Insert each tag
    for _, tag := range product.Tags {
        _, err := tx.Exec(context.Background(), stmt, productId,tag)
        if err != nil {
            // Rollback the transaction if an error occurs
            tx.Rollback(context.Background())
            return err
        }
    }
	err = tx.Commit(context.Background())

	if err!=nil{
		return err
	}

	return nil
}

func (pm *ProductModel) Update(user_id int, productId int, product types.ProductCreate)(bool, error){
	tx,err := pm.db.Begin(context.Background())

	if err!=nil{
		return false, nil
	}
	defer tx.Rollback(context.Background())


	// This might be dumbest approach KEKW

	result,err := tx.Exec(context.Background(), "UPDATE product SET name=$1, price=$2, imageUrl=$3, condition=$4, isPurchaseable=$5 WHERE id=$6 AND user_id=$7",
	product.Name,product.Price,product.ImageURL,product.Condition,product.IsPurchaseable,productId,user_id)

	if err != nil {
		return false,err
    }

	rowsAffected := result.RowsAffected()
    if rowsAffected == 0 {
        return false,nil
    }

	_,err = tx.Exec(context.Background(), "UPDATE product_tags SET is_deleted=true WHERE product_id=$1",
			productId)

	if err != nil {
		tx.Rollback(context.Background())
		return false, err
    }

	stmt:= "INSERT INTO product_tags (product_id, tag) VALUES ($1, $2)"
    // Insert each tag
    for _, tag := range product.Tags {
        _, err := tx.Exec(context.Background(), stmt, productId,tag)
        if err != nil {
            // Rollback the transaction if an error occurs
            tx.Rollback(context.Background())
            return false, err
        }
    }
	err = tx.Commit(context.Background())

	if err!=nil{
		return false, err
	}

	return true, nil
}

func (pm *ProductModel) Delete(user_id int, productId int)(bool, error){
	tx,err := pm.db.Begin(context.Background())

	if err!=nil{
		return false, nil
	}
	defer tx.Rollback(context.Background())


	result,err := tx.Exec(context.Background(), "UPDATE product SET is_deleted=true WHERE id=$1 AND user_id=$2",
						productId,user_id)

	if err != nil {
		return false,err
    }

	rowsAffected := result.RowsAffected()
    if rowsAffected == 0 {
        return false,nil
    }

	_,err = tx.Exec(context.Background(), "UPDATE product_tags SET is_deleted=true WHERE product_id=$1",
			productId)

	if err != nil {
		tx.Rollback(context.Background())
		return false, err
    }

	
	err = tx.Commit(context.Background())

	if err!=nil{
		return false, err
	}

	return true, nil
}

