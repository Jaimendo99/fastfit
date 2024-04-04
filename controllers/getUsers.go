package controllers

import (
	"fastfit/models"
	"fastfit/views"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User
		db.First(&user, c.Param("id"))
		return views.UserForList(user).Render(c.Request().Context(), c.Response().Writer)
	}
}

func GetUsers(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var users []models.User
		db.Find(&users)
		return views.UserList(users).Render(c.Request().Context(), c.Response().Writer)
	}
}
