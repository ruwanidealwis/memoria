package models

import (
	"fmt"
	"os"

	"memoria/app/utils"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() error {

	godotenv.Load()
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")
	name := os.Getenv("DB_NAME")
	host := os.Getenv("DB_HOST")

	dbUri := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", host, username, name, password)

	conn, err := gorm.Open(postgres.Open(dbUri), &gorm.Config{})

	if err != nil {
		return err
	}

	db = conn
	//db.Debug().AutoMigrate(&Scrapbook{}, &Page{}, &Image{}, &Song{}, &Map{})

	return nil
}

func GetDB() *gorm.DB {
	return db
}

func GetPageById(db *gorm.DB, id uint) (Page, error) {
	var page Page
	res := db.Model(&page).Where(`id = ?`, id).Preload("Song").Preload("Map").Preload("Images").First(&page)
	if res.Error != nil {
		return Page{}, res.Error
	}
	return page, nil
}

func GetImageById(db *gorm.DB, id uint) (Image, error) {
	var image Image
	res := db.Model(&image).Where(`id = ?`, id).First(&image)
	if res.Error != nil {
		return Image{}, res.Error
	}
	return image, nil
}

func CreatePage(db *gorm.DB,
	title string,
	images []Image,
	headingOne string,
	headingTwo string,
	headingThree string,
	location Map,
	song Song,
	scrapbookID uint) (Page, error) {
	page := Page{Title: title, Images: images, HeadingOne: headingOne, HeadingTwo: headingTwo, HeadingThree: headingThree, Map: location, Song: song, ScrapbookID: scrapbookID}
	res := db.Create(&page)
	if res.Error != nil {
		return Page{}, res.Error
	}
	return page, nil
}

func CreateScrapbook(db *gorm.DB, name string, userId uint) (Scrapbook, error) {
	scrapbook := Scrapbook{Name: name, UserID: userId}
	res := db.Create(&scrapbook)
	if res.Error != nil {
		return Scrapbook{}, res.Error
	}
	return scrapbook, nil
}

func GetScrapbookByName(db *gorm.DB, name string) (Scrapbook, error) {
	var sb Scrapbook
	res := db.Where("name = ?", name).First(&sb)
	if res.Error != nil {
		return Scrapbook{}, res.Error
	}
	return sb, nil
}

func CreateImage(db *gorm.DB, file []byte) (Image, error) {
	image := Image{File: utils.EncodeImage(file)}
	res := db.Create(&image)
	if res.Error != nil {
		return Image{}, res.Error
	}
	return image, nil
}

func SaveImage(db *gorm.DB, image string) (Image, error) {
	imgModel := Image{File: image}
	res := db.Create(&imgModel)
	if res.Error != nil {
		return Image{}, res.Error
	}
	return imgModel, nil
}

func CreateUser(db *gorm.DB, user User) (User, error) {
	res := db.Create(&user)
	if res.Error != nil {
		return User{}, res.Error
	}
	return user, nil
}

func GetUserByID(db *gorm.DB, id uint) (User, error) {
	user := User{}
	res := db.Where(`id = ?`, id).First(&user)
	if res.Error != nil {
		return User{}, res.Error
	}
	return user, nil
}
func FindUserByEmail(db *gorm.DB, email string) (User, error) {
	user := User{}
	fmt.Print(email);
	res := db.Where(`email = ?`, email).First(&user)

	if res.Error != nil {
		return User{}, res.Error
	}
	return user, nil
}
func GetScrapbooksByUser(db *gorm.DB, userId uint) ([]Scrapbook, error) {
	var scrapbooks []Scrapbook
	res := db.Where(`user_id = ?`, userId).Find(&scrapbooks)
	if res.Error != nil {
		return []Scrapbook{}, res.Error
	}
	return scrapbooks, nil
}
