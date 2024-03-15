package models

import (
	"context"
	"log"
	"strconv"

	"github.com/gabriel-tama/be-week-1/types"
	"github.com/gabriel-tama/be-week-1/types/response"
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

type FindAllProductParams struct {
	Limit          int
	Offset         int
	Tags           []string
	Condition      string
	ShowEmptyStock bool
	MaxPrice       float64
	MinPrice       float64
	SortBy         string
	OrderBy        string
	Search         string
}

type FindAllProductResponse struct {
	Products []response.ProductResponse
	Meta     types.Meta
}

func (pm *ProductModel) FindAll(param FindAllProductParams) (FindAllProductResponse, error) {
	var res FindAllProductResponse
	var products []response.ProductResponse
	var total int

	var query = `
	SELECT p.id, p.name, p.price, p.imageUrl, p.stock, p.condition, p.isPurchaseable, pt.tag, 
	(SELECT COUNT(*) FROM payment pp WHERE pp.product_id = p.id) AS purchaseCount
	FROM product p LEFT JOIN product_tags pt ON p.id = pt.product_id 
	WHERE p.is_deleted = false AND pt.is_deleted
	`

	if len(param.Tags) > 0 {
		query += "AND pt.tag IN ("
		for i, tag := range param.Tags {
			if i == 0 {
				query += "'" + tag + "'"
			} else {
				query += ", '" + tag + "'"
			}
		}
		query += ") "
	}

	if param.Condition != "" {
		query += "AND p.condition = '" + param.Condition + "' "
	}

	if !param.ShowEmptyStock {
		query += "AND p.stock > 0 "
	}

	if param.MaxPrice > 0 && param.MinPrice > 0 {
		maxPriceStr := strconv.FormatFloat(param.MaxPrice, 'f', -1, 64)
		minPriceStr := strconv.FormatFloat(param.MinPrice, 'f', -1, 64)
		query += "AND p.price BETWEEN " + minPriceStr + " AND " + maxPriceStr + " "

	} else {
		if param.MaxPrice > 0 {
			maxPriceStr := strconv.FormatFloat(param.MaxPrice, 'f', -1, 64)
			query += "AND p.price <= " + maxPriceStr + " "
		}

		if param.MinPrice > 0 {
			minPriceStr := strconv.FormatFloat(param.MinPrice, 'f', -1, 64)
			query += "AND p.price >= " + minPriceStr + " "
		}
	}

	if param.Search != "" {
		query += "AND p.name LIKE '%" + param.Search + "%' "
	}

	if param.SortBy != "" {
		query += "ORDER BY p." + param.SortBy + " "
		if param.OrderBy != "" {
			query += param.OrderBy + " "
		}
	}

	query += "LIMIT " + strconv.Itoa(param.Limit) + " OFFSET " + strconv.Itoa(param.Offset)

	rows, err := pm.db.Query(context.Background(), query)

	if err != nil {
		log.Fatal(err)
		return res, err
	}
	defer rows.Close()

	for rows.Next() {
		var product response.ProductResponse
		err := rows.Scan(
			&product.ProductId,
			&product.Name,
			&product.Price,
			&product.ImageURL,
			&product.Stock,
			&product.Condition,
			&product.IsPurchaseable,
			&product.Tags,
			&product.PurchaseCount,
		)

		if err != nil {
			return res, err
		}

		total++
		products = append(products, product)
	}
	if err := rows.Err(); err != nil {
		return res, err
	}

	res.Products = products
	res.Meta = types.Meta{
		Limit:  param.Limit,
		Offset: param.Offset,
		Total:  total,
	}

	return res, nil
}

func (pm *ProductModel) Create(user_id int, product types.ProductCreate) error {
	tx, err := pm.db.Begin(context.Background())

	if err != nil {
		return nil
	}
	defer tx.Rollback(context.Background())

	var productId int

	err = tx.QueryRow(context.Background(), "INSERT INTO product (user_id, name, price, imageUrl,stock,condition,isPurchaseable) VALUES ($1, $2, $3, $4,$5,$6,$7) RETURNING id",
		user_id, product.Name, product.Price, product.ImageURL, product.Stock, product.Condition, product.IsPurchaseable).Scan(&productId)

	if err != nil {
		return err
	}

	stmt := "INSERT INTO product_tags (product_id, tag) VALUES ($1, $2)"
	// Insert each tag
	for _, tag := range product.Tags {
		_, err := tx.Exec(context.Background(), stmt, productId, tag)
		if err != nil {
			// Rollback the transaction if an error occurs
			tx.Rollback(context.Background())
			return err
		}
	}
	err = tx.Commit(context.Background())

	if err != nil {
		return err
	}

	return nil
}

func (pm *ProductModel) Update(user_id int, productId int, product types.ProductCreate) (bool, error) {
	tx, err := pm.db.Begin(context.Background())

	if err != nil {
		return false, nil
	}
	defer tx.Rollback(context.Background())

	// This might be dumbest approach KEKW

	result, err := tx.Exec(context.Background(), "UPDATE product SET name=$1, price=$2, imageUrl=$3, condition=$4, isPurchaseable=$5 WHERE id=$6 AND user_id=$7",
		product.Name, product.Price, product.ImageURL, product.Condition, product.IsPurchaseable, productId, user_id)

	if err != nil {
		return false, err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return false, nil
	}

	_, err = tx.Exec(context.Background(), "UPDATE product_tags SET is_deleted=true WHERE product_id=$1",
		productId)

	if err != nil {
		tx.Rollback(context.Background())
		return false, err
	}

	stmt := "INSERT INTO product_tags (product_id, tag) VALUES ($1, $2)"
	// Insert each tag
	for _, tag := range product.Tags {
		_, err := tx.Exec(context.Background(), stmt, productId, tag)
		if err != nil {
			// Rollback the transaction if an error occurs
			tx.Rollback(context.Background())
			return false, err
		}
	}
	err = tx.Commit(context.Background())

	if err != nil {
		return false, err
	}

	return true, nil
}

func (pm *ProductModel) Delete(user_id int, productId int) (bool, error) {
	tx, err := pm.db.Begin(context.Background())

	if err != nil {
		return false, nil
	}
	defer tx.Rollback(context.Background())

	result, err := tx.Exec(context.Background(), "UPDATE product SET is_deleted=true WHERE id=$1 AND user_id=$2",
		productId, user_id)

	if err != nil {
		return false, err
	}

	rowsAffected := result.RowsAffected()
	if rowsAffected == 0 {
		return false, nil
	}

	_, err = tx.Exec(context.Background(), "UPDATE product_tags SET is_deleted=true WHERE product_id=$1",
		productId)

	if err != nil {
		tx.Rollback(context.Background())
		return false, err
	}

	err = tx.Commit(context.Background())

	if err != nil {
		return false, err
	}

	return true, nil
}
