package models

import (
	"time"

	"gorm.io/gorm"
)

type Invoices struct {
	gorm.Model

	ID            int            `json:"id" gorm:"primaryKey;autoIncrement" `
	InvoiceNumber string         `json:"invoicenumber" gorm:"size:100"`
	Date          time.Time      `json:"date"`
	Total         float64        `json:"total"`
	CustomerID    int            `json:"customer_id"`
	InvoiceLines  []InvoiceLines `gorm:"foreignKey:InvoiceID"`
}
