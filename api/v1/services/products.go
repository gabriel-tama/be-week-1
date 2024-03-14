package services

import (
	"github.com/gabriel-tama/be-week-1/models"
	"github.com/gabriel-tama/be-week-1/types"
)

type ProductService interface {
	Create(user_id int, product types.ProductCreate)(error)
}

type productServiceImpl struct {
    productModel *models.ProductModel
}

func NewProductService(productModel *models.ProductModel) ProductService {
    return &productServiceImpl{productModel: productModel}
}


func (ps *productServiceImpl) Create(user_id int, product types.ProductCreate)  (error) {
    err := ps.productModel.Create(user_id, product)
    if err != nil {
        return err
    }


    return nil
}
