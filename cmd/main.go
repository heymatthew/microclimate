package main

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type Sample struct {
	Path string
}

type Topography struct {
	Dir     string
	Samples []Sample
}

func (t *Topography) Load() error {
	return filepath.Walk(t.Dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		t.Samples = append(t.Samples, Sample{Path: path})
		return nil
	})
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html.tmpl", gin.H{
			"Visitor": "Doctor",
		})
	})
	router.Static("/static", "static")

	router.Run(":3000")
}
