package services

import (
	"github.com/gabriel-tama/be-week-1/models"
	"github.com/gabriel-tama/be-week-1/types"
)

type ProductService interface {
	Create(user_id int, product types.ProductCreate)(error)
	Update(user_id int, productId int, product types.ProductUpdate)(bool,error)
	Delete(user_id int, productId int)(bool,error)
	UpdateStock(user_id int, productId int, stock int)(bool,error)
	FindAll(props models.FindAllProductParams) (models.FindAllProductResponse, error)
	FindById(product_id int)(models.FindByIdResponse,error)
}

type productServiceImpl struct {
	productModel *models.ProductModel
}

func NewProductService(productModel *models.ProductModel) ProductService {
	return &productServiceImpl{productModel: productModel}
}

func (ps *productServiceImpl) FindAll(porps models.FindAllProductParams) (models.FindAllProductResponse, error) {
	res, err := ps.productModel.FindAll(porps)
	if err != nil {
		return models.FindAllProductResponse{}, err
	}

	return res, nil
}

func (ps *productServiceImpl) FindById(product_id int) (models.FindByIdResponse, error) {
	res, err := ps.productModel.FindById(product_id)
	if err != nil {
		return models.FindByIdResponse{}, err
	}

	return res, nil
}

func (ps *productServiceImpl) Create(user_id int, product types.ProductCreate) error {
	err := ps.productModel.Create(user_id, product)
	if err != nil {
		return err
	}

	return nil
}

func (ps *productServiceImpl) Update(user_id int, productId int, product types.ProductUpdate) (bool, error) {
	exist, err := ps.productModel.Update(user_id, productId, product)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (ps *productServiceImpl) Delete(user_id int, productId int) (bool, error) {
	exist, err := ps.productModel.Delete(user_id, productId)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (ps *productServiceImpl) UpdateStock(user_id int, productId int, stock int )(bool ,error){
	exist, err:= ps.productModel.UpdateStock(user_id,productId,stock)
	if err!=nil{
		return false, err
	}
	return exist,nil
}
