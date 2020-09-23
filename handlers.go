package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
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
	db, err := gorm.Open("sqlite3", "sample.db")
	if err != nil {
		panic("failed to connect database")
	}
	roomID := c.Param("room_id")
	t := Theater{}
	t.ID = roomID
	db.First(&t)
	db.Close()
	if t.MovieTitle == "" {
		c.Redirect(http.StatusBadRequest, "/")
	}
	// fmt.Println(t)
	render(c, gin.H{
		"payload": t,
	}, "theater.html")
}

func createTheater(c *gin.Context) {
	db, err := gorm.Open("sqlite3", "sample.db")
	if err != nil {
		panic("failed to connect database")
	}
	id, err := MakeRandomStr(10)
	if err != nil {
		panic(err)
	}

	//layout := "2006-01-02T15:04:05"
	start_datetime_str := c.Query("start_datetime") + ":00"
	//start_datetime, err := time.Parse(layout, start_datetime_str)
	if err != nil {
		panic(err)
	}
	t := &Theater{
		ID:            id,
		MovieTitle:    c.Query("movie_title"),
		MovieLink:     c.Query("movie_link"),
		StartDatetime: start_datetime_str,
	}

	db.Create(t)

	db.Close()
	var redirect_dst string
	redirect_dst = "//localhost:8080/theater/" + id
	c.Redirect(http.StatusSeeOther, redirect_dst)
}

func resetCounter(c *gin.Context) {
	println("RESET!!")
}

// Render one of HTML, JSON or CSV based on the 'Accept' header of the request
// If the header doesn't specify this, HTML is rendered, provided that
// the template name is present
func render(c *gin.Context, data gin.H, templateName string) {
	switch c.Request.Header.Get("Content-Type") {
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
