package controllers

import (
	"errors"
	"fmt"
	"memoria/app/models"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type Login struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password"  binding:"required"`
}

func (a *APIEnv) GetUserScrapbooks(id uint) ([]models.Scrapbook, error) {
	return models.GetScrapbooksByUser(a.DB, id)
}

//sign up
func (a *APIEnv) CreateUser(c *gin.Context) (models.User, error) {
	c.MultipartForm()
	user := models.User{}
	err := c.ShouldBindWith(&user, binding.FormMultipart)

	if err != nil {
		return models.User{}, err
	}
	existingUser, err := models.FindUserByEmail(a.DB, user.Email)
	if existingUser.ID != 0 {
		return models.User{}, errors.New("Account with this email already exists")
	}
	return models.CreateUser(a.DB, user)

}

func (a *APIEnv) GetUserByID(id uint) (models.User, error) {
	res, err := models.GetUserByID(a.DB, id)
	if err != nil {
		return models.User{}, err
	}
	return res, nil
}

//login
func (a *APIEnv) Login(c *gin.Context) (models.User, error) {

	loginDetails := Login{}
	err := c.BindJSON(&loginDetails)
	if err != nil {
		return models.User{}, err
	}
	fmt.Println(loginDetails.Email)
	user, err := models.FindUserByEmail(a.DB, loginDetails.Email)
	if err != nil {
		return models.User{}, err
	}
	if user.Password == loginDetails.Password {
		return user, nil
	} else {
		return models.User{}, errors.New("password/email mismatch")
	}

}
