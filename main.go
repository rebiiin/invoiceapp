package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rebiiin/invoiceapp/controllers"
	"github.com/rebiiin/invoiceapp/dbconfig"
	"github.com/rebiiin/invoiceapp/middlewares"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	dbconfig.ConnectDatabase()

	//dbconfig.DatabaseMigrate()

	router := gin.Default()
	api := router.Group("/api")

	//Users
	api.POST("/users/signup", controllers.SignupUser)
	api.POST("users/login", controllers.LoginUser)
	api.PUT("/users/changePassword/:id", middlewares.AuthJwt(), controllers.ChangePasswordUser)
	api.GET("/users/", middlewares.AuthJwt(), controllers.GetUsers)
	api.GET("/users/:id/", middlewares.AuthJwt(), controllers.GetUserById)
	api.GET("/users/delete/:id/", middlewares.AuthJwt(), controllers.DeleteUser)

	//Refresh Token
	api.POST("/RefreshToken", controllers.RefreshToken)

	//Customers
	api.GET("customers/", middlewares.AuthJwt(), controllers.GetCustomers)
	api.GET("customers/:id", middlewares.AuthJwt(), controllers.GetCustomerById)
	api.GET("customers/delete/:id", middlewares.AuthJwt(), controllers.DeleteCustomer)
	api.PUT("customers/update/:id", middlewares.AuthJwt(), controllers.UpdateCustomer)
	api.POST("customers/create/", middlewares.AuthJwt(), controllers.CreateCustomer)

	//Suppliers
	api.GET("suppliers/", middlewares.AuthJwt(), controllers.GetSuppliers)
	api.GET("suppliers/:id", middlewares.AuthJwt(), controllers.GetSupplierById)
	api.GET("suppliers/delete/:id", middlewares.AuthJwt(), controllers.DeleteSupplier)
	api.PUT("suppliers/update/:id", middlewares.AuthJwt(), controllers.UpdateSupplier)
	api.POST("suppliers/create/", middlewares.AuthJwt(), controllers.CreateSupplier)

	//Products
	api.GET("products/", middlewares.AuthJwt(), controllers.GetProducts)
	api.GET("products/:id", middlewares.AuthJwt(), controllers.GetProductById)
	api.GET("products/delete/:id", middlewares.AuthJwt(), controllers.DeleteProduct)
	api.PUT("products/update/:id", middlewares.AuthJwt(), controllers.UpdateProduct)
	api.PATCH("products/updateImage/:id", middlewares.AuthJwt(), controllers.UploadProductImage)
	api.POST("products/create/", middlewares.AuthJwt(), controllers.CreateProduct)

	//Invoices
	api.GET("invoices/", middlewares.AuthJwt(), controllers.GetInvoices)
	api.GET("invoices/:id", middlewares.AuthJwt(), controllers.GetInvoiceById)
	api.GET("invoices/delete/:id", middlewares.AuthJwt(), controllers.DeleteInvoice)
	api.POST("invoices/create/", middlewares.AuthJwt(), controllers.CreateInvoice)

	//Invoice lines
	api.GET("invoicelines/:invoiceId", middlewares.AuthJwt(), controllers.GetInvoiceLinesByInvoiceId)
	api.GET("invoicelines/delete/:id", middlewares.AuthJwt(), controllers.DeleteInvoiceLine)
	api.PUT("invoicelines/update/:id", middlewares.AuthJwt(), controllers.UpdateInvoiceLine)
	api.POST("invoicelines/create", middlewares.AuthJwt(), controllers.CreateInvoiceLine)

	log.Fatal(router.Run(":8080"))

}
