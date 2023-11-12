package middlewares

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/rebiiin/invoiceapp/helpers"
)

func AuthJwt() gin.HandlerFunc {

	return func(c *gin.Context) {
		bearerToken := c.GetHeader("Authorization")

		cleanedToken := strings.TrimPrefix(bearerToken, "Bearer ")

		if bearerToken == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "The authorization token is not found"})
			c.Abort()
			return
		}

		jwt := helpers.Jwt{}
		_, err := jwt.ValidateToken(cleanedToken)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid token"})
			c.Abort()
			return
		}

		c.Next()

	}

}
