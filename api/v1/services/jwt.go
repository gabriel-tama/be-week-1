package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)


type JWTService interface {
    CreateToken(username string) (string, error)
    ValidateToken(tokenString string) error
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

func (jwtService *jwtServiceImpl) ValidateToken(tokenString string) error {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return jwtService.secretKey, nil
    })

    if err != nil {
        return err
    }

    if !token.Valid {
        return jwt.ErrSignatureInvalid
    }

    return nil
}