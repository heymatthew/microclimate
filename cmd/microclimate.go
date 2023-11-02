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
	topo := SetupTopography()
	fmt.Println(topo)

	router := SetupRouter()
	err := router.Run(":3000")
	if err != nil {
		log.Fatal(err)
	}
}

func SetupTopography() pkg.Cache {
	topo := pkg.Cache{Dir: CacheDir()}
	err := topo.Load()
	if err != nil {
		log.Fatal(err)
	}
	return topo
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
