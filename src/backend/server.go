package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.LoadHTMLFiles("src/frontend/home/home.html")

	server.Static("/src", "./src")

	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"title": "Main website",
		})
	})

	server.GET("/file-corruptor", func(c *gin.Context) {
		c.String(http.StatusOK, "File corruptor page")
	})

	server.Run(":8080")
}
