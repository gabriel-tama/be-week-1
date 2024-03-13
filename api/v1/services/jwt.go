package services

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


type JWTService interface {
    CreateToken(username string) (string, error)
    ValidateToken(tokenString string) (*jwt.Token,error)
}

type jwtServiceImpl struct {
    secretKey []byte
}

func NewJWTService(secretKey string) JWTService {
    return &jwtServiceImpl{secretKey: []byte(secretKey)}
}

func (jwtService *jwtServiceImpl) CreateToken(username string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "username": username,
        "exp":      time.Now().Add(time.Hour * 24).Unix(),
    })

    tokenString, err := token.SignedString(jwtService.secretKey)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func (jwtService *jwtServiceImpl) ValidateToken(tokenString string) (*jwt.Token,error ){
  return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, isvalid := token.Method.(*jwt.SigningMethodHMAC); !isvalid {
			return nil, fmt.Errorf("Invalid token", token.Header["alg"])

		}
		return []byte(jwtService.secretKey), nil
	})
}