package models

import (
	"gorm.io/gorm"
)

type InvoiceLines struct {
	gorm.Model

	ID        int     `json:"id" gorm:"primaryKey;autoIncrement"`
	Quantity  int     `json:"quantity"`
	Price     float64 `json:"price"`
	ProductID int     `json:"product_id"`
	InvoiceID int     `json:"invoice_id"`
}
