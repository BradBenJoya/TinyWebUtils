package main

import (
    "log"
    "net/http"
    "os"
    "os/exec"
    "path/filepath"

    "github.com/gin-gonic/gin"
)

func main() {
    router := gin.Default()

    router.LoadHTMLFiles(
        "src/frontend/index.html",
        "src/frontend/corruptor/corruptor.html",
    )

    router.StaticFile("/corruptor/corruptor.js", "src/frontend/corruptor/api/corruptor.js")

    router.GET("/", func(c *gin.Context) {
        c.HTML(http.StatusOK, "index.html", gin.H{"title": "Main page"})
    })

    router.GET("/corruptor", func(c *gin.Context) {
        c.HTML(http.StatusOK, "corruptor.html", gin.H{"title": "Corruptor"})
    })

    router.POST("/corruptor/api", func(c *gin.Context) {
        log.Printf("upload request from %s", c.ClientIP())
        file, err := c.FormFile("files")
        if err != nil {
            log.Printf("upload rejected: could not read multipart file: %v", err)
            c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
            return
        }
        log.Printf("upload received: filename=%q size=%d content_type=%q", file.Filename, file.Size, file.Header.Get("Content-Type"))

        tmpDir, err := os.MkdirTemp("", "corruptor-*")
        if err != nil {
            log.Printf("failed to create temporary directory: %v", err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }
        defer os.RemoveAll(tmpDir)

        name := filepath.Base(file.Filename)
        inPath := filepath.Join(tmpDir, name)
        outPath := inPath + ".corrupted"

        if err := c.SaveUploadedFile(file, inPath); err != nil {
            log.Printf("failed to save uploaded file %q: %v", file.Filename, err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
            return
        }

        log.Printf("uploaded file saved: path=%q size=%d", inPath, file.Size)
        cmd := exec.Command("bin/corruptor/corruptor", inPath)
        log.Printf("starting corruptor: executable=%q input=%q", cmd.Path, inPath)
        if output, err := cmd.CombinedOutput(); err != nil {
            log.Printf("corruptor failed: %v; output: %s", err, string(output))
            c.JSON(http.StatusInternalServerError, gin.H{"error": "corruption failed"})
            return
        }

        info, err := os.Stat(outPath)
        if err != nil {
            log.Printf("corruptor output missing at %q: %v", outPath, err)
            c.JSON(http.StatusInternalServerError, gin.H{"error": "no output produced"})
            return
        }

        log.Printf("sending corrupted file: filename=%q size=%d", name+".corrupted", info.Size())
        c.FileAttachment(outPath, name+".corrupted")
    })

    if err := router.Run(":8000"); err != nil {
        log.Fatalf("server failed: %v", err)
    }
}