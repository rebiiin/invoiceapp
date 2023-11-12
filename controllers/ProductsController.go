package controllers

import (
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
	"github.com/rebiiin/invoiceapp/dbconfig"
	"github.com/rebiiin/invoiceapp/models"
)

func GetProducts(c *gin.Context) {
	var products []models.Products

	result := dbconfig.DB.Find(&products)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, gin.H{"Products": products})

}

func GetProductById(c *gin.Context) {

	id := c.Param("id")

	var product models.Products

	result := dbconfig.DB.First(&product, id)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, product)

}

func DeleteProduct(c *gin.Context) {

	id := c.Param("id")
	var product models.Products
	result := dbconfig.DB.First(&product, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Product not found"})
		c.Abort()
		return

	}

	dbconfig.DB.Where("id = ?", id).Delete(&product)
	c.JSON(http.StatusOK, gin.H{"Product:" + id: "has been deleted."})

}

func CreateProduct(c *gin.Context) {
	var newProduct models.Products

	if err := c.ShouldBindJSON(&newProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}

	result := dbconfig.DB.Create(&newProduct)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": result.Error.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "New product created successfully"})

}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var product models.Products

	result := dbconfig.DB.First(&product, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Product not found"})
		c.Abort()
		return

	}

	var requestProduct struct {
		Name     string  `json:"name"`
		Barcode  string  `json:"barcode"`
		Quantity int     `json:"quantity"`
		Price    float64 `json:"price"`

		SupplierID int `json:"supplier_id"`
	}

	if err := c.ShouldBindJSON(&requestProduct); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid input"})
		c.Abort()
		return
	}

	product.Name = requestProduct.Name
	product.Barcode = requestProduct.Barcode
	product.Quantity = requestProduct.Quantity
	product.Price = requestProduct.Price
	product.SupplierID = requestProduct.SupplierID

	dbconfig.DB.Save(&product)

	c.JSON(http.StatusOK, product)

}

func UploadProductImage(c *gin.Context) {

	id := c.Param("id")

	var product models.Products

	result := dbconfig.DB.First(&product, id)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{"Error": "Product image not found"})
		c.Abort()
		return

	}

	var requestProductImage struct {
		file *multipart.FileHeader `from:"file"`
	}

	if err := c.ShouldBind(&requestProductImage); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "Invalid input image file"})
		c.Abort()
		return
	}

	file, handler, err := c.Request.FormFile("file")

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": "No file provided"})
		c.Abort()
		return
	}
	defer file.Close()

	filename := filepath.Join("./uploadImages", handler.Filename)

	out, err := os.Create(filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}

	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": err.Error()})
		c.Abort()
		return
	}

	product.Image = filename

	if err := dbconfig.DB.Save(&product).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"Error": "Internal server error"})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, product)

}
