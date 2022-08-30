package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/colors", getColors)

	r.GET("/", func(c *gin.Context) {
		c.String(200, "tinyplm!")
	})

	r.Run(":9000")
}

func getColors(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, colors)
}

type color struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	RGB         string `json:"rgb"`
	Description string `json:"description"`
}

var colors = []color{
	{ID: "1", Name: "Blue Train", RGB: "John Coltrane", Description: "56.99"},
	{ID: "2", Name: "Jeru", RGB: "Gerry Mulligan", Description: "17.99"},
	{ID: "3", Name: "Sarah Vaughan and Clifford Brown", RGB: "Sarah Vaughan", Description: "39.99"},
}
