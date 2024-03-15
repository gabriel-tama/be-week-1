package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTService interface {
	CreateToken(user_id int) (string, error)
	ValidateToken(tokenString string) (*jwt.Token, error)
	GetUserIDByToken(tokenString string) (string, error)
}

type jwtServiceImpl struct {
	secretKey []byte
}

func NewJWTService(secretKey string) JWTService {
	return &jwtServiceImpl{secretKey: []byte(secretKey)}
}

func (jwtService *jwtServiceImpl) CreateToken(user_id int) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user_id,
		"exp":     time.Now().Add(time.Minute * 2).Unix(),
	})

	tokenString, err := token.SignedString(jwtService.secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (jwtService *jwtServiceImpl) ValidateToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("invalid token: %v", token.Header["alg"])
		}

		return []byte(jwtService.secretKey), nil
	})
}

func (jwtService *jwtServiceImpl) GetUserIDByToken(tokenString string) (string, error) {
	aToken, err := jwtService.ValidateToken(tokenString)
	if err != nil {
		// panic(err.Error())
		return "", err
	}
	claims := aToken.Claims.(jwt.MapClaims)
	return fmt.Sprintf("%v", claims["user_id"]), nil
}
