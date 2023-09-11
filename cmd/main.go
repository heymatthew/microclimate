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

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("web/templates/*")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html.tmpl", gin.H{
			"Visitor": "Doctor",
		})
	})
	router.Static("/static", "web/static")

	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(usr.HomeDir)

	topo := pkg.Topography{Dir: usr.HomeDir + "/.microclimate"}
	err = topo.Load()
	if err != nil {
		log.Fatal(err)
	}

	err = router.Run(":3000")
	if err != nil {
		log.Fatal(err)
	}
}
