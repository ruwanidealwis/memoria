package test

import (
	"fmt"
	"memoria/app/controllers"
	"memoria/app/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/suite"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type TestSuiteEnv struct {
	suite.Suite
	db *gorm.DB
}

type APIEnv struct {
	DB *gorm.DB
}

// Tests are run before they start
func (suite *TestSuiteEnv) SetupSuite() {
	fmt.Printf("hi hello goodbye")
	SetupTest()
	suite.db = models.GetDB()
}

// Running after each test
func (suite *TestSuiteEnv) TearDownTest() {
	ClearTable(suite.db)
}

// Running after all tests are completed
func (suite *TestSuiteEnv) TearDownSuite() {
	//suite.db.Close()
	ClearTable(suite.db)
}

// This gets run automatically by `go test`
func TestSuite(t *testing.T) {
	suite.Run(t, new(TestSuiteEnv))
}

//Setup router
func SetupTest() *gin.Engine {
	router := gin.Default()
	err := models.Connect()
	if err != nil {
		fmt.Println(err)
	}

	api := &controllers.APIEnv{
		DB: models.GetDB(),
	}

	// Initialize all routes
	controllers.InitializeRoutes(router, api)

	return router
}

func (suite *TestSuiteEnv) CreateTestUser(db *gorm.DB) (models.User, error) {
	user := models.User{
		FirstName: "John",
		LastName:  "Doe",
		Email:     "jdoe@email.com",
		Password:  "jdoe123!",
	}

	if err := db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func ClearTable(db *gorm.DB) {
	db.Exec("DELETE FROM page_images")
	db.Exec("DELETE FROM maps")
	db.Exec("DELETE FROM images")
	db.Exec("DELETE FROM songs")
	db.Exec("DELETE FROM pages")
	db.Exec("DELETE FROM scrapbooks")
	db.Exec("ALTER SEQUENCE scrapbooks_id_seq RESTART WITH 1")
}

// func (suite *TestSuiteEnv) Test_GetScrapbook_EmptyResult() {
// 	req, w := setGetScrapbookRouter(suite.db)

// 	a := assert.New(suite.T())
// 	a.Equal(http.MethodGet, req.Method, "HTTP request method error")
// 	a.Equal(http.StatusNotFound, w.Code, "HTTP request status code error")

// 	body, err := ioutil.ReadAll(w.Body)
// 	if err != nil {
// 		a.Error(err)
// 	}

// 	actual := models.Scrapbook{}
// 	if err := json.Unmarshal(body, &actual); err != nil {
// 		a.Error(err)
// 	}

// 	expected := models.Scrapbook{}
// 	a.Equal(expected, actual)
// }

func setGetScrapbookRouter(db *gorm.DB) (*http.Request, *httptest.ResponseRecorder) {
	r := gin.New()
	req, err := http.NewRequest(http.MethodGet, "/scrapbook", nil)
	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return req, w
}
