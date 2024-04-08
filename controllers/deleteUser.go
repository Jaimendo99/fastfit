package controllers

import (
	"fastfit/models"
	"fastfit/store"
	"net/http"

	"github.com/labstack/echo/v4"
)

func DeleteUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User
		store.DB.Delete(&user, c.Param("id"))
		return c.NoContent(http.StatusOK)
	}
}
