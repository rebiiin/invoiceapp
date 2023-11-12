package models

import (
	"gorm.io/gorm"
)

type Suppliers struct {
	gorm.Model

	ID       int        `json:"id" gorm:"primaryKey;autoIncrement"`
	Name     string     `json:"name" gorm:"size:50"`
	Phone    string     `json:"phone" gorm:"size:20"`
	Products []Products `json:"-" gorm:"foreignKey:SupplierID"`
}
