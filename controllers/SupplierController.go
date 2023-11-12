package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rebiiin/invoiceapp/dbconfig"
	"github.com/rebiiin/invoiceapp/models"
)

func GetSuppliers(c *gin.Context) {
	var suppliers []models.Suppliers

	result := dbconfig.DB.Find(&suppliers)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"suppliers": suppliers})

}

func GetSupplierById(c *gin.Context) {

	id := c.Param("id")

	var supplier models.Suppliers

	result := dbconfig.DB.First(&supplier, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, supplier)

}

func DeleteSupplier(c *gin.Context) {

	id := c.Param("id")
	var supplier models.Suppliers
	result := dbconfig.DB.First(&supplier, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Supplier not found"})
		c.Abort()
		return

	}

	dbconfig.DB.Where("id = ?", id).Delete(&supplier)
	c.JSON(http.StatusOK, gin.H{"Supplier:" + id: "has been deleted."})

}

func UpdateSupplier(c *gin.Context) {
	id := c.Param("id")

	var supplier models.Suppliers

	result := dbconfig.DB.First(&supplier, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Supplier not found"})
		c.Abort()
		return

	}

	var requestSupplier struct {
		Name  string `json:"name"`
		Phone string `json:"phone"`
	}

	if err := c.ShouldBindJSON(&requestSupplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid input"})
		c.Abort()
		return
	}

	supplier.Name = requestSupplier.Name
	supplier.Phone = requestSupplier.Phone

	dbconfig.DB.Save(&supplier)

	c.JSON(http.StatusOK, supplier)

}

func CreateSupplier(c *gin.Context) {
	var newSupplier models.Suppliers

	if err := c.ShouldBindJSON(&newSupplier); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}

	//inputs validation

	// errorValidation := helpers.InputValidation(&newSupplier)
	// if len(errorValidation) > 0 {
	// 	c.JSON(http.StatusBadRequest, gin.H{"Error": errorValidation})
	// 	return
	// }

	result := dbconfig.DB.Create(&newSupplier)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "New supplier created successfully"})

}
