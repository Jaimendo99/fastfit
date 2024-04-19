package controllers

import (
	"fastfit/models"
	"fastfit/store"
	"fastfit/views"

	"github.com/labstack/echo/v4"
)

func GetPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		var users []models.User
		store.DB.Find(&users)
		content := views.Index(users)
		return views.Layout(content, "Home", false).Render(c.Request().Context(), c.Response().Writer)
	}
}
