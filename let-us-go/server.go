package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/go-xorm/xorm"

	_ "github.com/go-sql-driver/mysql"

	"./controllers"
	"./modules"
)

type ServerConfig struct {
	Database struct{

	}

	Server struct{

	}
}

func initDatabase(driver, conntection string) (*xorm.Engine, error) {
	db, err := xorm.NewEngine(driver, conntection)
	if err != nil {
		return nil, err
	}

	if err = db.Sync(new(modules.Test)); err != nil {
		panic(err)
	}
	return db, nil
}

func main() {
	db, err := initDatabase("mysql", "root:123456@tcp(127.0.0.1:3306)/hdd?charset=utf8")
	if err != nil {
		panic(err)
	}
	defer db.Close()

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