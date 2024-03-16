package middlewares

import (
	"log"
	"net/http"

	"github.com/gabriel-tama/be-week-1/api/v1/services"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwtService services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "unauthorized"})
			return
		}

		const BEARER_SCHEMA = "BEARER "
		tokenString := authHeader[len(BEARER_SCHEMA):]
		_, err := jwtService.ValidateToken(tokenString)

		if err != nil {
			log.Println(err)
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message": "forbidden"})
		}
	}
}
