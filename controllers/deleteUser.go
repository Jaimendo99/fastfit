package controllers

import (
	"fastfit/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func DeleteUser(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var user models.User
		db.Delete(&user, c.Param("id"))
		return c.NoContent(http.StatusOK)
	}
}
