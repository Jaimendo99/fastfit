package main

import (
	"fmt"

	"fastfit/controllers"
	"fastfit/store"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {

	e := echo.New()
	e.Use(middleware.Logger())

	db, err := store.NewConnection("test.db")

	if err != nil {
		e.Logger.Fatal(err)
		fmt.Printf("error opening sqlite:" + err.Error())
		return
	}

	store.Migrate(db)

	e.GET("/", controllers.GetPage(db), middleware.Logger())

	e.GET("/users/:id", controllers.GetUser(db), middleware.Logger())

	e.GET("/users", controllers.GetUsers(db), middleware.Logger())

	e.POST("/users", controllers.CreateUser(db), middleware.Logger())
	e.GET("/userForm", controllers.GetAddUserForm(), middleware.Logger())

	e.DELETE("/users/:id", controllers.DeleteUser(db), middleware.Logger())

	e.PATCH("/users/:id", controllers.UpdateUser(db), middleware.Logger())
	e.GET("userForm/:id", controllers.GetUpdateForm(db), middleware.Logger())

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":8080"))
}
