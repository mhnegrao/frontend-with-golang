package main

import (
	"embed"
	"net/http"
	"path"

	"github.com/gin-gonic/gin"
)

var (
	buildDir = "frontend/build"

	//go:embed all:frontend/build
	f embed.FS
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		index := path.Join(buildDir, "index.html")
		if b, err := f.ReadFile(index); err == nil {
			c.Writer.Write(b)
		}
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	// route for svelte build files
	// after-request middleware
	r.Use(func(c *gin.Context) {
		trypath := path.Join(buildDir, c.Request.URL.Path)
		if _, err := f.ReadFile(trypath); err == nil {
			c.Status(http.StatusOK)
			c.FileFromFS(trypath, http.FS(f))
			return
		}

		index := path.Join(buildDir, "index.html")
		if b, err := f.ReadFile(index); err == nil {
			c.Status(http.StatusOK)
			c.Header("Content-Type", "text/html")
			c.Writer.Write(b)
		}
	})

	r.Run("127.0.0.1:8080")
}
