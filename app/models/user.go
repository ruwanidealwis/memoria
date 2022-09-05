package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	FirstName string      `json:"firstName" form:"firstName" binding:"required"`
	LastName  string      `json:"lastName" form:"lastName" binding:"required"`
	Email     string      `json:"email"  form:"email"  binding:"required"`
	Password  string      `json:"password" form:"password" binding:"required" `
	Scrapbook []Scrapbook `json:"pages"`
}
