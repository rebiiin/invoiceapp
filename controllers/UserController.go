package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/rebiiin/invoiceapp/dbconfig"
	"github.com/rebiiin/invoiceapp/helpers"
	"github.com/rebiiin/invoiceapp/models"
)

// Register user
func SignupUser(c *gin.Context) {

	var user models.User
	var existingUser models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}

	//inputs validation
	errorValidation := helpers.InputValidation(user)
	if len(errorValidation) > 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": errorValidation})
		return
	}

	dbconfig.DB.Find(&existingUser, "email = ?", user.Email)

	if existingUser.ID != 0 {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "The email address is already in use!"})
		c.Abort()
		return
	}

	//Hashing user password
	hashedPassword, err := helpers.HashPassword(user.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error hashing password": err.Error()})
		c.Abort()
		return
	}

	user.Password = hashedPassword

	result := dbconfig.DB.Create(&user)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully", "Id": user.ID, "email": user.Email})

}

func validateStruct(user models.User) {
	panic("unimplemented")
}

//Login user

func LoginUser(c *gin.Context) {
	var userlogin models.User

	if err := c.ShouldBindJSON(&userlogin); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid input"})
		c.Abort()
		return
	}

	var foundUser models.User

	result := dbconfig.DB.Where("email = ?", userlogin.Email).First(&foundUser)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid email"})
		c.Abort()
		return
	}

	if err := helpers.VerifyPassword(foundUser.Password, userlogin.Password); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"Error": "Invalid password"})
		c.Abort()
		return
	}

	jwt := helpers.Jwt{}
	token, err := jwt.CreateToken(foundUser)

	if err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Unable to create Access Token"})
		c.Abort()
		return

	}

	c.JSON(http.StatusOK, gin.H{"message": "Login successful", "token": token})

}

//LoginOut user

func LoginOutUser(c *gin.Context) {

}

//Get users

func GetUsers(c *gin.Context) {
	var users []models.User

	result := dbconfig.DB.Find(&users)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})

}

// Get User By Id
func GetUserById(c *gin.Context) {

	id := c.Param("id")

	var user models.User

	result := dbconfig.DB.First(&user, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, user)

}

// Change password
func ChangePasswordUser(c *gin.Context) {

	id := c.Params.ByName("id")

	var user models.User

	if err := dbconfig.DB.First(&user, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "User not found"})
		c.Abort()
		return
	}

	var newPassword struct {
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&newPassword); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid input"})
		c.Abort()
		return
	}

	//Hashing user password
	hashedPassword, err := helpers.HashPassword(newPassword.Password)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error hashing password": err.Error()})
		c.Abort()
		return
	}

	user.Password = hashedPassword
	dbconfig.DB.Save(&user)

	c.JSON(http.StatusOK, user)
}

//Delete user

func DeleteUser(c *gin.Context) {

	id := c.Params.ByName("id")
	var user models.User
	result := dbconfig.DB.First(&user, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "User not found"})
		c.Abort()
		return

	}

	dbconfig.DB.Where("id = ?", id).Delete(&user)
	c.JSON(http.StatusOK, gin.H{"User:" + id: "has been deleted."})

}
