package main

import (
	"github.com/labstack/echo/v4"
	"github.com/timileyin19/echo_framework/router"
	"github.com/timileyin19/echo_framework/server"
)

func main() {
	// initialize echo 
	app := echo.New()


	// load all routers
	router.LoadAllRouters(app)

	// listen and serve on port 3000
	server.SetServer(app)

}
