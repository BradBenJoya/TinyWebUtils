package main

import (
    "fmt"
    "net/http"
    "path/filepath"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.LoadHTMLFiles(
        "src/frontend/index.html",
        "src/frontend/corruptor/corruptor.html",
    )

    router.StaticFile("/corruptor/api", "src/frontend/corruptor/api/corruptor.js")

    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{"title": "Main page"})
    })

    router.GET("/corruptor", func(c *gin.Context) {
        c.HTML(http.StatusOK, "corruptor.html", gin.H{"title": "Corruptor"})
    })

    router.POST("/corruptor/api", func(c *gin.Context) {
        file, err := c.FormFile("files")
        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }

        f, err := file.Open()
        if err != nil {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        defer f.Close()

        contentType := file.Header.Get("Content-Type")
        if contentType == "" {
            contentType = "application/octet-stream"
        }

        c.DataFromReader(
            http.StatusOK,
            file.Size,
            contentType,
            f,
            map[string]string{
                "Content-Disposition": `attachment; filename="` + filepath.Base(file.Filename) + `"`,
            },
        )
    })

    router.Run(":8000")
}