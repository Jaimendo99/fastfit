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

	// userData := models.User{Name: "John Doe", Username: "johndoe", Password: "password", Email: "jaime@gmail.com"}

	// result := db.Create(&userData)

	// if result.Error != nil {
	// 	e.Logger.Fatal(result.Error)
	// 	fmt.Printf("error creating user:" + result.Error.Error())
	// 	return
	// }

	// fmt.Println("User created with ID:" + strconv.FormatUint(uint64(userData.ID), 10))

	e.GET("/users", controllers.GetUsers(db), middleware.Logger())

	e.Logger.Fatal(e.Start(":8080"))
}
