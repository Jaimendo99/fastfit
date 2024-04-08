package controllers

import (
	"fastfit/models"
	"fastfit/store"
	"fastfit/views"

	"github.com/labstack/echo/v4"
)

func UpdateUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User
		store.DB.First(&user, c.Param("id"))
		user.Name = c.FormValue("name")
		user.Username = c.FormValue("username")
		user.Password = c.FormValue("password")
		user.Email = c.FormValue("email")
		store.DB.Save(&user)
		return views.UserForList(user).Render(c.Request().Context(), c.Response().Writer)
	}
}

func GetUpdateForm() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User
		store.DB.First(&user, c.Param("id"))
		return views.UpdateForm(user).Render(c.Request().Context(), c.Response().Writer)
	}
}
