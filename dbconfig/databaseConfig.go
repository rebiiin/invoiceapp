package dbconfig

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/rebiiin/invoiceapp/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbUser := os.Getenv("DBUSER")
	dbPassword := os.Getenv("DBPASSWORD")
	dbHost := os.Getenv("DBHOST")
	dbPort := os.Getenv("DBPORT")
	dbName := os.Getenv("DBNAME")

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
		panic("Failed to connect to database!")
	}
	log.Println("Connected to Database!")

}

func DatabaseMigrate() {
	DB.AutoMigrate(&models.Customers{}, &models.InvoiceLines{}, &models.Invoices{}, &models.Products{}, &models.Suppliers{}, &models.User{})
	log.Println("Database Migration Completed!")
}
