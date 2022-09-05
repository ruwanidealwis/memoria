package controllers

import (
	"memoria/app/models"

	"github.com/gin-gonic/gin"
)

type CreateRequest struct {
	Name   string `json:"name" binding:"required"`
	UserID uint   `json:"user" binding:"required"`
}

func (a *APIEnv) CreateScrapbook(c *gin.Context) (models.Scrapbook, error) {
	body := CreateRequest{}
	if err := c.BindJSON(&body); err != nil {
		return models.Scrapbook{}, err
	} else {
		res, dberr := models.CreateScrapbook(a.DB, body.Name, body.UserID)
		if dberr != nil {
			return models.Scrapbook{}, dberr
		}
		return res, nil

	}
}

func (a *APIEnv) GetScrapbook(name string) (models.Scrapbook, error) {
	res, err := models.GetScrapbookByName(a.DB, name)
	if err != nil {
		return models.Scrapbook{}, err
	}
	return res, nil
}
