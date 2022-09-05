package models

import (
	"gorm.io/gorm"
)

type Scrapbook struct {
	gorm.Model
	Page   []Page `json:"pages"`
	Name   string `json:"name"`
	UserID uint   `json:"user"`
}
