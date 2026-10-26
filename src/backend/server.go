package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.LoadHTMLFiles("src/frontend/home/home.html", "src/frontend/corruptor/corruptor.html")

	server.Static("/src", "./src")

	server.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{
			"title": "Main website",
		})
	})

	server.GET("/file-corruptor", func(c *gin.Context) {
		c.HTML(http.StatusOK, "corruptor.html", gin.H{
			"title": "File corruptor",
		})
	})

	server.POST("/file-corruptor/api", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Failed to retrieve form files",
			})
			return
		}

		files := form.File["file"]

		if len(files) == 0 {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "No files uploaded",
			})
			return
		}

		tmpDir, err := os.MkdirTemp("", "tmp-*")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create temporary directory",
			})
			return
		}
		defer os.RemoveAll(tmpDir)

		for _, file := range files {
			dst := filepath.Join(tmpDir, filepath.Base(file.Filename))

			if err := c.SaveUploadedFile(file, dst); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Failed to save uploaded file to disk",
				})
				return
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"message": "Files uploaded successfully",
			"count":   len(files),
		})
	})

	server.Run(":8080")

	server.Run(":8080")
}
