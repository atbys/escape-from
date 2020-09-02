package main

func initializeRoutes() {
	router.GET("/", showIndexPage)
	router.GET("/theater/:room_num", showTheaterPage)
	router.GET("/create_room", createTheater)
}
