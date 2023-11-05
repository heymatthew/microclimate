package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os/user"

	"github.com/gin-gonic/gin"
	"github.com/heymatthew/microclimate/pkg"
	"github.com/heymatthew/microclimate/web"
)

type Sample struct {
	Path string
}

func CacheDir() string {
	usr, err := user.Current()
	if err != nil {
		fmt.Printf("Cannot find current user: %v\n", err)
		return ""
	}
	return usr.HomeDir + "/.microclimate"
}

func main() {
	cache := SetupCache()
	fmt.Println(cache)

	router := SetupRouter()
	err := router.Run(":3000")
	if err != nil {
		log.Fatal(err)
	}
}

func SetupCache() pkg.Cache {
	cache := pkg.Cache{Dir: CacheDir()}
	err := cache.Load()
	if err != nil {
		log.Fatal(err)
	}
	return cache
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	templ := template.Must(template.New("").ParseFS(web.Files, "templates/*.tmpl"))
	router.SetHTMLTemplate(templ)

	router.StaticFS("/static", http.FS(web.Static))

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html.tmpl", gin.H{
			"title": "Microclimate index",
		})
	})
	return router
}
