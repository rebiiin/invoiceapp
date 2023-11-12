package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/rebiiin/invoiceapp/dbconfig"
	"github.com/rebiiin/invoiceapp/models"
)

func GetInvoices(c *gin.Context) {

	var invoiceData []models.Invoices

	if err := dbconfig.DB.Preload("InvoiceLines").Find(&invoiceData).Error; err != nil {

		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"Invoices": invoiceData})
}

func GetInvoiceById(c *gin.Context) {

	id := c.Param("id")

	var invoiceData models.Invoices

	result := dbconfig.DB.Preload("InvoiceLines").First(&invoiceData, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, invoiceData)

}

func DeleteInvoice(c *gin.Context) {

	id := c.Param("id")
	var invoiceData models.Invoices
	result := dbconfig.DB.First(&invoiceData, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Invoice not found"})
		c.Abort()
		return

	}

	dbconfig.DB.Where("id = ?", id).Delete(&invoiceData)

	c.JSON(http.StatusOK, gin.H{"Invoices:" + id: "has been deleted."})

}

func CreateInvoice(c *gin.Context) {

	var newInvoice models.Invoices

	var customer models.Customers

	if err := c.ShouldBindJSON(&newInvoice); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}
	invocieTotal := newInvoice.Total
	customerID := newInvoice.CustomerID

	result := dbconfig.DB.First(&customer, customerID)
	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "CustomerID not found"})
		c.Abort()
		return

	}

	// when customer blance is not enough!
	if customer.Balance < invocieTotal {

		c.JSON(http.StatusBadRequest, gin.H{"Error": "The customer balance is less than the invoice total! please add balance for customer : " + customer.FirstName + " " + customer.LastName})
		c.Abort()
		return
	}

	// Update the customer's balance
	customer.Balance -= invocieTotal

	// Generate Invoice Number-ID

	invoiceYear := strconv.Itoa(newInvoice.Date.Year())

	var countRows int64

	resultCount := dbconfig.DB.Model(&newInvoice).Count(&countRows)
	if resultCount.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": resultCount.Error.Error()})
		c.Abort()
		return
	}

	newInvoice.InvoiceNumber = invoiceYear + "-" + fmt.Sprintf("%03d", countRows+1)

	result = dbconfig.DB.Save(&customer)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error.Error()})
		c.Abort()
		return
	}

	// Begin transaction
	txn := dbconfig.DB.Begin()

	if err := txn.Create(&newInvoice).Error; err != nil {

		//Role back if faild
		txn.Rollback()

		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}

	// Commit transaction
	txn.Commit()

	c.JSON(http.StatusCreated, gin.H{"message": "New invoice created successfully"})

}
