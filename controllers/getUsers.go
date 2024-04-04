package controllers

import (
	"fastfit/models"
	"fastfit/views"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetUsers(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var users []models.User
		db.Find(&users)
		return views.Index(users).Render(c.Request().Context(), c.Response().Writer)
	}
}
