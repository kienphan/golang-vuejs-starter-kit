package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var host = "localhost"

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("./*.html")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", "")
	})

	router.Run(":3000")
}
