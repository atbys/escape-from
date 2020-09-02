package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func showIndexPage(c *gin.Context) {
	dsc := "Welcome to virtual theater"
	render(c, gin.H{
		"title":       "Home Page",
		"description": dsc,
	},
		"index.html",
	)
}

func showTheaterPage(c *gin.Context) {
	render(c, gin.H{
		"room_number": c.Param("room_num"),
		"endtime":     "2020/09/02 18:00:00",
	}, "theater.html")
}

func createTheater(c *gin.Context){
	render(c, gin.H{
		"title": 
	}, "theater.html")
}
// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {

	switch c.Request.Header.Get("Accept") {
	case "application/json":
		// Respond with JSON
		c.JSON(http.StatusOK, data["payload"])
	case "application/xml":
		// Respond with XML
		c.XML(http.StatusOK, data["payload"])
	default:
		// Respond with HTML
		c.HTML(http.StatusOK, templateName, data)
	}

}
