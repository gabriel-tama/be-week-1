package types

type ProductCreate struct {
	Name           string   `json:"name" binding:"required,min=5,max=60"`
	Price          float64  `json:"price" binding:"required,min=0"`
	ImageURL       string   `json:"imageUrl" binding:"required,url"`
	Stock          int      `json:"stock" binding:"required,min=0"`
	Condition      string   `json:"condition" binding:"required,oneof=new second"`
	Tags           []string `json:"tags" binding:"required"`
	IsPurchaseable bool     `json:"isPurchaseable" binding:"required"`
}

type ProductStock struct {
	Stock int `json:"stock" binding:"required,min=0"`
}