package controllers

import (
	"io/ioutil"
	"memoria/app/models"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"gorm.io/gorm"
)

// Remove dependency on DB object
type APIEnv struct {
	DB *gorm.DB
}
type FormData struct {
	HeadingOne   string `form:"h1" binding:"required"`
	HeadingTwo   string `form:"h2" binding:"required"`
	HeadingThree string `form:"h3" binding:"required"`
	SpotifyID    string `form:"song" binding:"required"`
	Location     string `form:"location" binding:"required"`
	Title        string `form:"page-title" binding:"required"`
	ScrapbookID  uint   `form:"scrapbook-id" binding:"required"`
	MapImage     string `form:"maps" binding:"required"`
}

func (a *APIEnv) CreatePage(c *gin.Context) (models.Page, error) {
	formFiles, _ := c.MultipartForm()
	//c.MultipartForm()

	var form FormData
	err := c.ShouldBindWith(&form, binding.FormMultipart)
	if err != nil {
		return models.Page{}, err
	}

	song := models.Song{SpotifyID: form.SpotifyID}

	if err != nil {
		return models.Page{}, err
	}

	mapCover, err := models.SaveImage(a.DB, form.MapImage)
	if err != nil {
		return models.Page{}, err
	}
	maps := models.Map{Location: form.Location, ImageID: mapCover.ID}

	images := []models.Image{}
	//convert file uploads to bytes
	for i := 1; i <= 3; i++ {
		imageInput, err := formFiles.File["image-"+strconv.Itoa(i)][0].Open()
		if err != nil {
			return models.Page{}, err
		}
		imageBytes, err := ioutil.ReadAll(imageInput)
		if err != nil {
			return models.Page{}, err
		}
		image, err := models.CreateImage(a.DB, imageBytes)
		images = append(images, image)
		if err != nil {
			return models.Page{}, err
		}
	}

	res, err := models.CreatePage(a.DB, form.Title, images, form.HeadingOne, form.HeadingTwo, form.HeadingThree, maps, song, form.ScrapbookID)
	if err != nil {
		return models.Page{}, err
	}
	return res, nil
}

func (a *APIEnv) GetPage(id uint) (models.Page, error) {
	res, err := models.GetPageById(a.DB, id)
	if err != nil {
		return models.Page{}, err
	}
	return res, nil
}

func (a *APIEnv) GetImageById(id uint) (models.Image, error) {
	res, err := models.GetImageById(a.DB, id)
	if err != nil {
		return models.Image{}, err
	}
	return res, nil
}
