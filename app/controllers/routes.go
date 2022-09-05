package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getLanding(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"landing.html",
		gin.H{
			"title": "Memoria",
		},
	)
}

func getLoginPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"login.html",
		gin.H{
			"title": "Login",
		},
	)
}

func getSignUpPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"signUp.html",
		gin.H{
			"title": "Create an account",
		},
	)
}

func getCreatePage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"create.html",
		gin.H{
			"title": "Create",
		},
	)
}

func getScrapbooksPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"dashboard.html",
		gin.H{
			"title": "My Scrapbooks",
		},
	)
}

func getExportPage(c *gin.Context) {
	c.HTML(
		http.StatusOK,
		"exportPage.html",
		gin.H{
			"title": "My Scrapbook Page",
		},
	)
}

func (a *APIEnv) getPage(c *gin.Context) {
	pageID := c.Param("id")
	id, err := strconv.ParseUint(pageID, 10, 32)
	if err != nil {
		c.JSON(500, "Something went wrong")
	} else {
		res, err := a.GetPage(uint(id))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}

func (a *APIEnv) createPage(c *gin.Context) {
	res, err := a.CreatePage(c)
	if err != nil {
		fmt.Print(err.Error())
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)

}

func (a *APIEnv) getScrapbook(c *gin.Context) {
	sbName := c.Query("name")
	res, err := a.GetScrapbook(sbName)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, res)
}

func (a *APIEnv) createScrapbook(c *gin.Context) {
	res, err := a.CreateScrapbook(c)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusCreated, res)

}

func (a *APIEnv) GetImage(c *gin.Context) {

	imageID := c.Query("id")
	id, err := strconv.ParseUint(imageID, 10, 32)
	if err != nil {
		c.JSON(500, "Something went wrong")
	} else {
		res, err := a.GetImageById(uint(id))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, res)
	}

}

func SearchSongs(c *gin.Context) {
	song := c.Query("name")
	songs, err := SearchSong(song)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, songs)
}

func getSong(c *gin.Context) {
	songID := c.Query("id")
	song, err := GetSongById(songID)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, song)
}

func getMap(c *gin.Context) {
	location := c.Query("location")
	googleMap, err := GetMapByLocation(location)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, googleMap)
}

func (a *APIEnv) SignUp(c *gin.Context) {
	user, err := a.CreateUser(c)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (a *APIEnv) UserLogin(c *gin.Context) {
	user, err := a.Login(c)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (a *APIEnv) GetScrapbooksUser(c *gin.Context) {
	userID := c.Param("userID")
	id, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}
	scrapbooks, err := a.GetUserScrapbooks(uint(id))

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, scrapbooks)
}
func (a *APIEnv) GetUser(c *gin.Context) {
	userID := c.Param("id")
	id, err := strconv.ParseUint(userID, 10, 32)
	if err != nil {
		c.JSON(500, "Bad ID")
	} else {
		res, err := a.GetUserByID(uint(id))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(http.StatusOK, res)
	}
}
func InitializeRoutes(router *gin.Engine, api *APIEnv) {
	router.GET("/", getLanding)
	router.GET("/sign-up", getSignUpPage)
	router.GET("/login", getLoginPage)
	router.GET("/scrapbooks", getScrapbooksPage)
	router.GET("/export", getExportPage)
	router.GET("/create", getCreatePage)
	router.GET("/search/song", SearchSongs)
	router.GET("/song", getSong)
	router.POST("/scrapbook", api.createScrapbook)
	router.GET("/scrapbook", api.getScrapbook)
	router.POST("/page", api.createPage)
	router.GET("/page/:id", api.getPage)
	router.GET("/image", api.GetImage)
	router.GET("/map", getMap)
	router.GET("/scrapbooks/:userID", api.GetScrapbooksUser)
	router.POST("user/sign-up", api.SignUp)
	router.POST("user/login", api.UserLogin)
	router.GET("user/:id", api.GetUser)

}
