package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Idea struct {
	gorm.Model
	User string `gorm:"unique_index:idx_user_idea"`
	Idea string
}

var db *gorm.DB

func main() {
	var err error
	db, err = gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("Failed to open database")
	}

	db.AutoMigrate(&Idea{})

	router := gin.Default()

	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})

	router.Static("/static", "./static")

	router.POST("/submit", func(c *gin.Context) {
		user := c.PostForm("user")
		idea := c.PostForm("idea")

		// Verificar si el usuario ya ha enviado una idea
		var existingIdea Idea
		if db.Where("user = ?", user).First(&existingIdea).Error == nil {
			c.HTML(http.StatusOK, "index.html", gin.H{"err": "Ya has enviado una idea"})
			return
		}

		// Insertar la nueva idea
		db.Create(&Idea{User: user, Idea: idea})
		c.Redirect(http.StatusFound, "/ideas")
	})

	router.GET("/ideas", func(c *gin.Context) {
		var ideas []Idea
		db.Find(&ideas)

		c.HTML(http.StatusOK, "ideas.html", gin.H{
			"ideas": ideas,
		})
	})

	router.Run(":8080")
}
