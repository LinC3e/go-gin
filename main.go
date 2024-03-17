package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "home",
		})
	})

	r.GET("/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title": "| News List",
		})
	})

	r.POST("/add", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "add",
		})
	})

	r.PUT("/edit", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "edit",
		})
	})

	r.DELETE("/delete", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "delete",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
