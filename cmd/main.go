package main

import (
	"fmt"
	"net/http"
	"os"

	"fastfit/controllers"
	"fastfit/inits"
	"fastfit/store"

	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type jwtUserClaims struct {
	Sub uint `json:"sub"`
	jwt.RegisteredClaims
}

func init() {
	inits.LoadEnvVariables()

	store.InitDB()

	store.Migrate(store.DB)
}

func main() {

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	config := echojwt.Config{
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return &jwtUserClaims{}
		},

		SigningKey: []byte(os.Getenv("JWT_SECRET")),

		TokenLookup: "cookie:Authorization",

		ErrorHandler: func(c echo.Context, err error) error {
			fmt.Printf("JWT Error: %v\n", err) // Log the error for debugging

			// if err.Error() == "missing value in cookies" {
			// 	return c.JSON(http.StatusUnauthorized, echo.Map{"error": "missing token"})
			// }

			// if strings.Contains(err.Error(), "signature is invalid") {
			// 	return c.JSON(http.StatusUnauthorized, echo.Map{"error": "invalid token"})
			// }

			// if strings.Contains(err.Error(), "expired") {
			// 	fmt.Printf("Token expired\n")
			// 	return c.Redirect(301, "/login")
			// }

			return c.Redirect(http.StatusPermanentRedirect, "/login")
		},
	}

	e.GET("/login", controllers.GetLogin(), middleware.Logger())
	e.POST("/login", controllers.SignIn(), middleware.Logger())
	e.POST("/signup", controllers.SignUp(), middleware.Logger())

	r := e.Group("/", middleware.Logger())

	r.Use(echojwt.WithConfig(config))

	r.GET("users/:id", controllers.GetUser(), middleware.Logger())

	r.GET("users", controllers.GetUsers(), middleware.Logger())
	e.POST("/users", controllers.CreateUser(), middleware.Logger())

	r.GET("userForm", controllers.GetAddUserForm(), middleware.Logger())

	e.DELETE("/users/:id", controllers.DeleteUser(), middleware.Logger(), echojwt.WithConfig(config))

	r.PATCH("users/:id", controllers.UpdateUser(), middleware.Logger())
	r.GET("userForm/:id", controllers.GetUpdateForm(), middleware.Logger())

	e.Static("/static", "static")
	r.GET("", controllers.GetPage(), middleware.Logger())

	e.Logger.Fatal(e.Start(":" + os.Getenv("PORT")))
}
