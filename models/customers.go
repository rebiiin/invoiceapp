package models

import (
	"gorm.io/gorm"
)

type Customers struct {
	gorm.Model

	ID        int        `json:"id" gorm:"autoIncrement;primary_key" `
	FirstName string     `json:"firstname" gorm:"size:50"`
	LastName  string     `json:"lastNname" gorm:"size:50"`
	Address   string     `json:"address" gorm:"size:150"`
	Phone     string     `json:"phone" gorm:"size:20"`
	Balance   float64    `json:"balance"`
	Invoices  []Invoices `gorm:"foreignKey:CustomerID"`
}
