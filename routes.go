package main

func initializeRoutes() {
	router.GET("/", showIndexPage)
	router.GET("/theater/:room_id", showTheaterPage)
	router.GET("/create_room", createTheater)
}
