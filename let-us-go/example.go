package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"

	"controllers"	
)

func main() {
	server := echo.New()

	server_header := &http.Server{
		Addr:         ":1323",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	controllers.HomeController{}.Init(server.Group("/"))

	server.Logger.Fatal(server.StartServer(server_header))
}
