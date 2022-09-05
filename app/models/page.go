package models

import (
	"gorm.io/gorm"
)

type Page struct {
	gorm.Model
	Title        string  `json:"title"`
	Images       []Image `json:"images" gorm:"many2many:page_images;"`
	HeadingOne   string  `json:"headingOne"`
	HeadingTwo   string  `json:"headingTwo"`
	HeadingThree string  `json:"headingThree" `
	Song         Song    `json:"song"`
	Map          Map     `json:"map"`
	ScrapbookID  uint    `json:"scrapbook"`
}
