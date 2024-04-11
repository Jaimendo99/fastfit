package controllers

import (
	"fastfit/views"

	"github.com/labstack/echo/v4"
)

func GetLogin() echo.HandlerFunc {
	return func(c echo.Context) error {
		login := views.Login()
		return views.Layout(login, "Login", true).Render(c.Request().Context(), c.Response().Writer)
	}
}
