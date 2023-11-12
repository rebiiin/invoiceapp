package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rebiiin/invoiceapp/helpers"
	"github.com/rebiiin/invoiceapp/models"
)

func RefreshToken(c *gin.Context) {

	requestToken := models.Token{}

	if err := c.ShouldBindJSON(&requestToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid input"})
		c.Abort()
		return
	}

	jwt := helpers.Jwt{}
	user, err := jwt.ValidateRefreshToken(requestToken)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid token"})
		c.Abort()
		return
	}

	newToken, err := jwt.CreateToken(user)
	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Unable to create Access Token"})
		c.Abort()
		return

	}

	c.JSON(http.StatusOK, gin.H{"token": newToken})

}
