package types

type LoginData struct {
	Username string `json:"username" binding:"required,min=5,max=15"`
	Password string `json:"password" binding:"required,min=5,max=15"`
}
