package main

func initializeRoutes() {
	router.GET("/", showIndexPage)
	theater := router.Group("/theater")
	{
		theater.GET("/create", createTheater)
		theater.GET("/:room_num", showTheaterPage)
	}

}
