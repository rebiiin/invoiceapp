package models

import "gorm.io/gorm"

type Products struct {
	gorm.Model

	ID           int            `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string         `json:"name" gorm:"size:50"`
	Barcode      string         `json:"barcode" gorm:"size:50"`
	Quantity     int            `json:"quantity"`
	Price        float64        `json:"price"`
	Image        string         `json:"image" gorm:"size:250"`
	SupplierID   int            `json:"supplier_id"`
	InvoiceLines []InvoiceLines `json:"-" gorm:"foreignKey:ProductID"`
}
