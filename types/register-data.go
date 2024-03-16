package types

type RegisterData struct {
	Username string `json:"username" binding:"required,min=5,max=15"`
	Name     string `json:"name" binding:"required,min=5,max=50"`
	Password string `json:"password" binding:"required,min=5,max=15"`
}
