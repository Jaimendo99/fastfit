package controllers

import (
	"fastfit/models"
	"fastfit/views"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func UpdateUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User
		db.First(&user, c.Param("id"))
		user.Name = c.FormValue("name")
		user.Username = c.FormValue("username")
		user.Password = c.FormValue("password")
		user.Email = c.FormValue("email")
		db.Save(&user)
		return views.UserForList(user).Render(c.Request().Context(), c.Response().Writer)
	}
}

func GetUpdateForm(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User
		db.First(&user, c.Param("id"))
		return views.UpdateForm(user).Render(c.Request().Context(), c.Response().Writer)
	}
}
