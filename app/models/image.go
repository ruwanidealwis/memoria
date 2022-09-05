package models

import (
	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	File string `json:"file"`
	Map  []Map  `json:"maps"`
}
