package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Nombre   string `json:"Nombre" binding:"required"`
	Email    string `json:"Email" gorm:"unique" binding:"required,email"`
	Password string `json:"-" binding:"required"`
}
