package main

import (
	"os"

	"fastfit/controllers"
	"fastfit/inits"
	"fastfit/store"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	inits.LoadEnvVariables()

	store.InitDB()

	store.Migrate(store.DB)
}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/login", controllers.GetLogin(), middleware.Logger())
	e.POST("/login", controllers.SignIn(), middleware.Logger())
	e.POST("/signup", controllers.SignUp(), middleware.Logger())

	e.GET("/", controllers.GetPage(), middleware.Logger())

	e.GET("/users/:id", controllers.GetUser(), middleware.Logger())

	e.GET("/users", controllers.GetUsers(), middleware.Logger())

	e.POST("/users", controllers.CreateUser(), middleware.Logger())
	e.GET("/userForm", controllers.GetAddUserForm(), middleware.Logger())

	e.DELETE("/users/:id", controllers.DeleteUser(), middleware.Logger())

	e.PATCH("/users/:id", controllers.UpdateUser(), middleware.Logger())
	e.GET("userForm/:id", controllers.GetUpdateForm(), middleware.Logger())

	e.Static("/static", "static")

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
