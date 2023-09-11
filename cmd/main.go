package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Sample struct {
	Path string
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html.tmpl", gin.H{
			"Visitor": "Doctor",
		})
	})
	router.Static("/static", "web")

	err := router.Run(":3000")
	if err != nil {
		fmt.Println(fmt.Errorf("Failed to start: %v", err))
	}
}
