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

	// _, err = tx.Exec(context.Background(), "INSERT INTO product (user_id, name, price, imageUrl,stock,condition,isPurchaseable,tags) VALUES ($1, $2, $3, $4,$5,$6,$7,$8)",
	// user_id,product.Name,product.Price,product.ImageURL,product.Stock,product.Condition,product.IsPurchaseable,product.Tags)

	if err != nil {
		return err
    }
	err = tx.Commit(context.Background())

	if err!=nil{
		return err
	}

	return nil
}


