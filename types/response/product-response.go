package response

import "github.com/gabriel-tama/be-week-1/types"

type ProductResponse struct {
	ProductId      string   `json:"productId"`
	Name           string   `json:"name"`
	Price          float64  `json:"price"`
	ImageURL       string   `json:"imageUrl"`
	Stock          int      `json:"stock"`
	Condition      string   `json:"condition"`
	Tags           []string `json:"tags"`
	IsPurchaseable bool     `json:"isPurchaseable"`
	PurchaseCount  int      `json:"purchaseCount"`
}

type SellerResponse struct {
	Name             string
	ProductSoldTotal int
	BankAccounts     []types.GetBankData
}


