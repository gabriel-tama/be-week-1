package middlewares

import (
	"net/http"

	"github.com/gabriel-tama/be-week-1/api/v1/services"
	"github.com/gin-gonic/gin"
)

func AuthorizeJWT(jwtService services.JWTService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		const BEARER_SCHEMA = "BEARER "
		tokenString:= authHeader[len(BEARER_SCHEMA):]
		if authHeader == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message":"invalid request"})
			return
		}
		token, _ := jwtService.ValidateToken(tokenString)
	
		if !token.Valid{
			c.AbortWithStatusJSON(http.StatusForbidden, gin.H{"message":"forbidden"})
		}
	}
}