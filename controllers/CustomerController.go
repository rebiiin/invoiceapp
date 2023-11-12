package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rebiiin/invoiceapp/dbconfig"
	"github.com/rebiiin/invoiceapp/models"
)

func GetCustomers(c *gin.Context) {
	var customers []models.Customers

	result := dbconfig.DB.Find(&customers)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"customers": customers})

}

func GetCustomerById(c *gin.Context) {

	id := c.Param("id")

	var customer models.Customers

	result := dbconfig.DB.First(&customer, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, customer)

}

func DeleteCustomer(c *gin.Context) {

	id := c.Param("id")
	var customer models.Customers
	result := dbconfig.DB.First(&customer, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Customer not found"})
		c.Abort()
		return

	}

	dbconfig.DB.Where("id = ?", id).Delete(&customer)
	c.JSON(http.StatusOK, gin.H{"Customer:" + id: "has been deleted."})

}

func UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	var customer models.Customers
	result := dbconfig.DB.First(&customer, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Customer not found"})
		c.Abort()
		return

	}

	var requestCustomer struct {
		FirstName string  `json:"firstname"`
		LastName  string  `json:"lastname"`
		Address   string  `json:"address"`
		Phone     string  `json:"phone"`
		Balance   float64 `json:"balance"`
	}

	if err := c.ShouldBindJSON(&requestCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid input"})
		c.Abort()
		return
	}

	customer.FirstName = requestCustomer.FirstName
	customer.LastName = requestCustomer.LastName
	customer.Address = requestCustomer.Address
	customer.Phone = requestCustomer.Phone
	customer.Balance = requestCustomer.Balance

	dbconfig.DB.Save(&customer)

	c.JSON(http.StatusOK, customer)

}

func CreateCustomer(c *gin.Context) {
	var newCustomer models.Customers

	if err := c.ShouldBindJSON(&newCustomer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}

	//inputs validation

	// errorValidation := helpers.InputValidation(&newCustomer)
	// if len(errorValidation) > 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{"Error": errorValidation})
	// 	return
	// }

	result := dbconfig.DB.Create(&newCustomer)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "New customer created successfully"})

}
