package types

type PaymentCreate struct {
	BankAccountId        string `json:"bankAccountId" binding:"required"`
	PaymentProofImageUrl string `json:"paymentProofImageUrl" binding:"required,url"`
	Quantity             int    `json:"quantity" binding:"required,min=1"`
}