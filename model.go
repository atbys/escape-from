package main

type Theater struct {
	ID            string `json:id`
	MovieTitle    string `json:movie_title`
	MovieLink     string `json:movie_link`
	StartDatetime string `json:start_datetime`
}
