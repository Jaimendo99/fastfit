package controllers

import (
	"fastfit/models"
	"fastfit/views"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func CreateUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := models.User{
			Name:     c.FormValue("name"),
			Username: c.FormValue("username"),
			Password: c.FormValue("password"),
			Email:    c.FormValue("email"),
		}

		result := db.Create(&user)
		if result.Error != nil {
			c.Response().WriteHeader(http.StatusInternalServerError)
			return views.Errors("Error when creating user").Render(c.Request().Context(), c.Response().Writer)
		}
		c.Response().WriteHeader(http.StatusCreated)
		return views.UserForList(user).Render(c.Request().Context(), c.Response().Writer)
	}
}

func GetAddUserForm() echo.HandlerFunc {
	return func(c echo.Context) error {
		return views.UserForm().Render(c.Request().Context(), c.Response().Writer)
	}
}
