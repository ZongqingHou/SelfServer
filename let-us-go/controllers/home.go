package controllers

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo"
)

type HomeController struct {
}

func (controller HomeController) Init(group *echo.Group) {
	group.GET("", controller.Get)
	group.GET("hello", controller.Controller)
}

func (HomeController) Get(context echo.Context) error {
	serverRequest := context.Request()
	contextPtr := serverRequest.Context()
	fmt.Println(contextPtr)
	return context.String(http.StatusOK, "Hello, World!")
}

func (HomeController) Controller(context echo.Context) error {
	return context.String(http.StatusOK, "wow")
}