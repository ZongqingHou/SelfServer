package controllers

import (
	"net/http"

	"github.com/labstack/echo"
)

type HomeController struct {
}

func (controller HomeController) Init(group *echo.Group) {
	group.GET("", controller.Index)
	group.GET("login", controller.HomeLogin)
	group.GET("logon", controller.HomeLogon)
}

func (HomeController) Index(context echo.Context) error {
	return context.String(http.StatusOK, "Hello, World!")
}

func (HomeController) HomeLogin(context echo.Context) error {

	return context.String(http.StatusOK, "wow")
}

func (HomeController) HomeLogon(context echo.Context) error {
	return context.String(http.StatusOK, "wow")
}
