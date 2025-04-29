package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"./routes"
	"./utils"
)

var r = gin.Default()

func main() {

	// make index.html from frontend the default page with "" endpoint
	r.Static("/static", "./frontend")
	r.LoadHTMLFiles("./frontend/index.html")
	
	// add a route to serve the `index.html`
	r.GET("", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", nil)
	})
	
	// register routes
	routes.SetupRoutes(r)
	log.Fatal(r.Run(":8080"))
}