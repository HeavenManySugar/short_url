package main

import (
	"crypto/md5"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type ShortURL struct {
	gorm.Model
	Hash string `json:"hash" gorm:"index;uniqueIndex"`
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
		// Generate a unique hash
		var hash string
		var exists bool
		for {
			hash = fmt.Sprintf("%x", md5.Sum([]byte(fmt.Sprintf("%s-%d", shortURL.Url, time.Now().UnixNano()))))[:8]

			var count int64
			db.Model(&ShortURL{}).Where("hash = ?", hash).Count(&count)
			exists = count > 0

			if !exists {
				break
			}
			time.Sleep(1 * time.Millisecond)
		}

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

	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		var shortURLs []ShortURL
		db.Find(&shortURLs)
		c.HTML(http.StatusOK, "index.html", gin.H{
			"shortURLs": shortURLs,
		})
	})

	r.GET("/shortener.js", func(c *gin.Context) {
		c.File("templates/shortener.js")
	})

	r.Run(":3010") // listen and serve on :8080
}
