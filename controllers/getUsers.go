package controllers

import (
	"fastfit/models"
	"fastfit/store"
	"fastfit/views"

	"github.com/labstack/echo/v4"
)

func GetUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User
		store.DB.First(&user, c.Param("id"))
		return views.UserForList(user).Render(c.Request().Context(), c.Response().Writer)
	}
}

func GetUsers() echo.HandlerFunc {
	return func(c echo.Context) error {
		var users []models.User
		store.DB.Find(&users)
		return views.UserList(users).Render(c.Request().Context(), c.Response().Writer)
	}
}
