package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rebiiin/invoiceapp/dbconfig"
	"github.com/rebiiin/invoiceapp/models"
)

func GetInvoiceLinesByInvoiceId(c *gin.Context) {

	id := c.Param("invoiceId")

	var invoiceLines []models.InvoiceLines

	result := dbconfig.DB.Where("invoice_id  = ?", id).First(&invoiceLines)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, invoiceLines)

}

func DeleteInvoiceLine(c *gin.Context) {

	id := c.Param("id")
	var invoiceLine models.InvoiceLines
	result := dbconfig.DB.First(&invoiceLine, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Invoice line not found"})
		c.Abort()
		return

	}

	dbconfig.DB.Where("id = ?", id).Delete(&invoiceLine)
	c.JSON(http.StatusOK, gin.H{"Invoice line:" + id: "has been deleted."})

}

func CreateInvoiceLine(c *gin.Context) {
	var newInvoiceLine models.InvoiceLines

	if err := c.ShouldBindJSON(&newInvoiceLine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}

	result := dbconfig.DB.Create(&newInvoiceLine)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "New Invoice line created successfully"})

}

func UpdateInvoiceLine(c *gin.Context) {
	id := c.Param("id")

	var invoiceLine models.InvoiceLines

	result := dbconfig.DB.First(&invoiceLine, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Invoice line not found"})
		c.Abort()
		return

	}

	var requestInvoiceLine struct {
		Quantity  int     `json:"quantity"`
		Price     float64 `json:"price"`
		ProductID int     `json:"product_id"`
		InvoiceID int     `json:"invoice_id"`
	}

	if err := c.ShouldBindJSON(&requestInvoiceLine); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid input"})
		c.Abort()
		return
	}

	invoiceLine.Quantity = requestInvoiceLine.Quantity
	invoiceLine.Price = requestInvoiceLine.Price
	invoiceLine.ProductID = requestInvoiceLine.ProductID
	invoiceLine.InvoiceID = requestInvoiceLine.InvoiceID

	dbconfig.DB.Save(&invoiceLine)

	c.JSON(http.StatusOK, invoiceLine)

}
