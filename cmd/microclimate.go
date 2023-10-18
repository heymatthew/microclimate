package main

import (
	"fmt"
	"log"
	"net/http"
	"os/user"

	"github.com/gin-gonic/gin"
	"github.com/heymatthew/microclimate/pkg"
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

	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*")
	router.Static("/static", "web/static")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html.tmpl", gin.H{
			"Visitor": "Doctor",
		})
	})
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
