package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"

	"./controllers"
	// "./modules"
	// "./utils"
)

func main() {
	// message := utils.InitMessageProcesser()
	// utils.MessageClose(message)

	server := echo.New()

	server_header := &http.Server{
		Addr:         ":80",
		ReadTimeout:  20 * time.Minute,
		WriteTimeout: 20 * time.Minute,
	}

	controllers.HomeController{}.Init(server.Group("/"))
	// controllers.TestController{}.Init(server.Group("/test"))
	controllers.WechatController{}.Init(server.Group("/wechat"))

	// server.Use(utils.ContextSession(message))
	// server.Static("/static", "static")

	/*
		need to use the middleware for pre and use functions
	*/

	server.Logger.Fatal(server.StartServer(server_header))
}
