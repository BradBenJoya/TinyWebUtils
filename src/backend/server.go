package main

import (
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/gin-gonic/gin"
)

func corruptorPath() string {
	executable := "corruptor"
	if runtime.GOOS == "windows" {
		executable += ".exe"
	}
	return filepath.Join("bin", "corruptor", executable)
}

func main() {
	gin.SetMode(gin.ReleaseMode)
	server := gin.Default()

	server.LoadHTMLFiles(
		"src/frontend/home/home.html",
		"src/frontend/corruptor/corruptor.html",
	)

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
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "No files uploaded",
			})
			return
		}

		tmpDir, err := os.MkdirTemp("", "tmp-*")
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to create temporary directory to store uploaded file",
			})
			return
		}
		defer os.RemoveAll(tmpDir)

		filename := filepath.Base(file.Filename)
		uploadedPath := filepath.Join(tmpDir, filename)
		if err := c.SaveUploadedFile(file, uploadedPath); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to save uploaded file to disk",
			})
			return
		}

		command := exec.CommandContext(c.Request.Context(), corruptorPath(), uploadedPath)
		output, err := command.CombinedOutput()
		if err != nil {
			log.Printf("Corruptor failed for uploaded file %q: %v: %s", filename, err, output)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to corrupt uploaded file",
			})
			return
		}

		corruptedPath := uploadedPath + ".corrupted"
		if _, err := os.Stat(corruptedPath); err != nil {
			log.Printf("Corruptor did not create output for uploaded file %q: %v (output: %s)", filename, err, output)
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Corruptor did not produce an output file",
			})
			return
		}

		c.FileAttachment(corruptedPath, filename+".corrupted")
	})

	if err := server.Run(":8000"); err != nil {
		log.Fatal(err)
	}
}