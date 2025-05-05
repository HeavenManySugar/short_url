package main

import (
	"crypto/md5"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ShortURL struct {
	gorm.Model
	Hash string `json:"hash"`
	Url  string `json:"url"`
}

type ShortURLRequest struct {
	Url string `json:"url"`
}

type ShortURLResponse struct {
	Hash string `json:"hash"`
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	db.AutoMigrate(&ShortURL{})

	r := gin.Default()

	r.POST("/shorten", func(c *gin.Context) {
		var shortURL ShortURLRequest
		if err := c.ShouldBindJSON(&shortURL); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		hash := fmt.Sprintf("%x", md5.Sum([]byte(shortURL.Url)))[:8] // Generate an 8-character hash
		db.Create(&ShortURL{
			Url:  shortURL.Url,
			Hash: hash,
		})
		c.JSON(http.StatusOK, ShortURLResponse{Hash: hash})
	})

	r.GET("/:hash", func(c *gin.Context) {
		hash := c.Param("hash")
		var shortURL ShortURL
		if err := db.Where("hash = ?", hash).First(&shortURL).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "URL not found"})
			return
		}
		c.Redirect(http.StatusFound, "http://"+shortURL.Url)
	})

	r.Run(":3010") // listen and serve on :8080
}
