package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type TestController struct {
}

func (controller TestController) Init(group *echo.Group) {
	group.GET("", controller.Get)
	group.GET("/hello", controller.Controller)
}

func (TestController) Get(context echo.Context) error {
	return context.String(http.StatusOK, "Hello, World!")
}

func (TestController) Controller(context echo.Context) error {
	return context.String(http.StatusOK, "wow")
}