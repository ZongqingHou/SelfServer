package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"

	"./controllers"
	// "./modules"

	"./utils"
)

type ServerConfig struct {
	Database struct{

	}

	Server struct{

	}
}


func main() {
	db := utils.InitMySQL("mysql", "root:123456@tcp(127.0.0.1:3306)/hdd?charset=utf8")
	defer db.Close()

	rd := utils.InitRedis("localhost:6379", "", 0)
	defer rd.Close()

	server := echo.New()

	server_header := &http.Server{
		Addr:         ":1323",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	controllers.HomeController{}.Init(server.Group("/"))
	controllers.TestController{}.Init(server.Group("/test"))

	server.Static("/static", "static")

	/*
		need to use the middleware for pre and use functions
	*/


	server.Logger.Fatal(server.StartServer(server_header))
}