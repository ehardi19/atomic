package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// Registrant is a struct that holds registrant table
type Registrant struct {
	ID          uint   `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phonenumber"`
	JobTitle    string `json:"jobtitle"`
	Company     string `json:"company"`
	Time        string `json:"time"`
	ClassID     uint   `json:"classid"`
}

// Class is a struct that holds class table
type Class struct {
	ID            uint   `json:"id"`
	Topic         uint   `json:"topic"`
	Title         string `json:"title"`
	Date          string `json:"date"`
	Location      string `json:"location"`
	Price         uint   `json:"price"`
	Seat          uint   `json:"seat"`
	Desc          string `json:"desc"`
	Require       string `json:"require"`
	Tutor         string `json:"tutor"`
	TutorJob      string `json:"tutorjob"`
	TutorLinkedin string `json:"tutorlinkedin"`
}

var db *gorm.DB
var err error

func main() {

	db, err = gorm.Open("sqlite3", "./gorm.db")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()

	db.AutoMigrate(&Registrant{}, &Class{})

	r := gin.Default()

	r.GET("/registrant/", GetRegistrants)
	r.GET("/registrant/:id", GetRegistrant)
	r.POST("/registrant", CreateRegistrant)
	r.PUT("/registrant/:id", UpdateRegistrant)
	r.DELETE("/registrant/:id", DeleteRegistrant)

	r.GET("/class/", GetClasses)
	r.GET("/topic/:topic", GetClassesByType)
	r.GET("/class/:id", GetClass)
	r.POST("/class", CreateClass)
	r.PUT("/class/:id", UpdateClass)
	r.DELETE("/class/:id", DeleteClass)
	r.GET("/searchclass/:title", SearchClass)

	r.Run(":8080")
}

//// REGISTRANT CRUD

// CreateRegistrant is a function to create a new registrant
func CreateRegistrant(c *gin.Context) {

	var registrant Registrant
	c.BindJSON(&registrant)

	if err := db.Create(&registrant).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, registrant)
	}
}

// GetRegistrant is a function to get a registrant by id
func GetRegistrant(c *gin.Context) {

	id := c.Params.ByName("id")
	var registrant Registrant
	if err := db.Where("id = ?", id).First(&registrant).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, registrant)
	}
}

// GetRegistrants is a function to get all registrants
func GetRegistrants(c *gin.Context) {

	var registrants []Registrant
	if err := db.Find(&registrants).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, registrants)
	}
}

// UpdateRegistrant is a function to update registrant data by id
func UpdateRegistrant(c *gin.Context) {

	var registrant Registrant
	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).First(&registrant).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&registrant)

	db.Save(&registrant)
	c.JSON(200, registrant)
}

// DeleteRegistrant is a function to delete registrant by id
func DeleteRegistrant(c *gin.Context) {

	id := c.Params.ByName("id")
	var registrant Registrant
	d := db.Where("id = ?", id).Delete(&registrant)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

//// CLASS CRUD

// CreateClass is a create a new class
func CreateClass(c *gin.Context) {

	var class Class
	c.BindJSON(&class)

	if err := db.Create(&class).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, class)
	}
}

// GetClass is a function to get a class by id
func GetClass(c *gin.Context) {

	id := c.Params.ByName("id")
	var class Class
	if err := db.Where("id = ?", id).First(&class).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, class)
	}
}

// GetClasses is a function to get all classes
func GetClasses(c *gin.Context) {

	var classes []Class
	if err := db.Find(&classes).Order("date desc").Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, classes)
	}
}

// GetClassesByType is a function to get classes by topic
// 1 = Software Engineer
// 2 = Product Designer
// 3 = Businees Analyst
func GetClassesByType(c *gin.Context) {

	topic := c.Params.ByName("topic")
	var classes []Class
	if err := db.Where("topic = ?", topic).Find(&classes).Order("date desc").Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, classes)
	}
}

// UpdateClass is a function to update a class by id
func UpdateClass(c *gin.Context) {

	var class Class
	id := c.Params.ByName("id")

	if err := db.Where("id = ?", id).First(&class).Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	}
	c.BindJSON(&class)

	db.Save(&class)
	c.JSON(200, class)
}

// DeleteClass is a delete a class by id
func DeleteClass(c *gin.Context) {

	id := c.Params.ByName("id")
	var class Class
	d := db.Where("id = ?", id).Delete(&class)
	fmt.Println(d)
	c.JSON(200, gin.H{"id #" + id: "deleted"})
}

// SearchClass is a function to searching a class with similiar title
func SearchClass(c *gin.Context) {

	title := c.Params.ByName("title")
	var classes []Class
	title = "%" + title + "%"

	if err := db.Where("id LIKE ?", "title").Find(&classes).Order("date desc").Error; err != nil {
		c.AbortWithStatus(404)
		fmt.Println(err)
	} else {
		c.JSON(200, classes)
	}
}
