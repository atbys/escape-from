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

type Theater struct {
	ID            string `json:id`
	MovieTitle    string `json:movie_title`
	MovieLink     string `json:movie_link`
	StartDatetime string `json:start_datetime`
}

func showTheaterPage(c *gin.Context) {
	db, err := gorm.Open("sqlite3", "sample.db")
	if err != nil {
		panic("failed to connect database")
	}
	roomID := c.Param("room_num")
	t := Theater{}
	t.ID = roomID
	db.First(&t)
	db.Close()
	if t.MovieTitle == "" {
		c.Redirect(http.StatusBadRequest, "/")
	}
	render(c, gin.H{
		"roomNumber": roomID,
		"MovieTitle": t.MovieTitle,
		"MovieLink":  t.MovieLink,
		"startTime":  t.StartDatetime,
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

	t := &Theater{
		ID:            id,
		MovieTitle:    c.Query("movie_title"),
		MovieLink:     c.Query("movie_link"),
		StartDatetime: c.Query("start_time"),
	}

	db.Create(t)

	db.Close()
	var redirect_dst string
	redirect_dst = "//localhost:8080/theater/" + id
	c.Redirect(http.StatusSeeOther, redirect_dst)
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
