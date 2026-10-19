package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.LoadHTMLFiles(
		"src/frontend/index.html",
		"src/frontend/corruptor/corruptor.html",
	)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Main page",
		})
	})

	router.GET("/corruptor", func(c *gin.Context) {
		c.HTML(http.StatusOK, "corruptor.html", gin.H{
			"title": "Corruptor",
		})
	})

	router.Run(":8000")
}
