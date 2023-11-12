package models

import "gorm.io/gorm"

// type User struct {
// 	gorm.Model
// 	ID       int    `json:"id" gorm:"autoIncrement;primary_key;"`
// 	Email    string `json:"email" gorm:"size:150;unique" validate:"required,email"`
// 	Password string `json:"password" validate:"required"`
// }

type User struct {
	gorm.Model

	ID       int    `json:"id" gorm:"primaryKey:autoIncrement"`
	Email    string `json:"email" gorm:"unique;size:150" validate:"required,email"`
	Password string `json:"password" gorm:"size:500" validate:"required"`
}
