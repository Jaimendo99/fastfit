package controllers

import (
	"errors"
	"fastfit/models"
	"fastfit/store"
	"net/http"
	"net/mail"
	"os"
	"time"
	"unicode"

	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

type body struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func SignUp() echo.HandlerFunc {
	return func(c echo.Context) error {

		var body body

		if c.Bind(&body) != nil {
			return c.JSON(http.StatusBadRequest, "Invalid request body")
		}

		err := validateData(&body)

		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}

		hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Error hashing password")
		}

		user := models.User{
			Name:     body.Name,
			Email:    body.Email,
			Username: body.Username,
			Password: string(hash),
		}

		result := store.DB.Create(&user)

		if result.Error != nil {
			return c.JSON(http.StatusInternalServerError, "Error creating user")
		}

		return c.JSON(http.StatusCreated, user)

	}

}

func SignIn() echo.HandlerFunc {
	return func(c echo.Context) error {
		var body struct {
			Username string `json:"username" form:"username"`
			Email    string `json:"email" form:"email"`
			Password string `json:"password" form:"password"`
		}

		if c.Bind(&body) != nil {
			return c.JSON(http.StatusBadRequest, "Invalid request body")
		}

		var user models.User

		if body.Username != "" {
			store.DB.Where("username = ?", body.Username).First(&user)
		} else {
			store.DB.Where("email = ?", body.Email).First(&user)
		}

		if user.ID == 0 {
			return c.JSON(http.StatusUnauthorized, "User not found")
		}

		err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

		if err != nil {
			return c.JSON(http.StatusUnauthorized, "Invalid password")
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": user.ID,
			"exp": time.Now().Add(time.Hour * 24 * 7).Unix(),
		})

		tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

		if err != nil {
			return c.JSON(http.StatusInternalServerError, "Error generating token")
		}

		c.SetCookie(&http.Cookie{
			Name:     "Authorization",
			Value:    tokenString,
			Expires:  time.Now().Add(time.Hour * 24 * 7),
			HttpOnly: true,
			Secure:   true,
		})

		return c.JSON(http.StatusOK, user)

	}
}

func validateData(body *body) error {
	if len(body.Name) < 4 {
		return errors.New("name is required, and must be at least 4 characters")
	}

	if len(body.Username) < 3 {
		return errors.New("username is required, and must be at least 3 characters")
	}

	_, err := mail.ParseAddress(body.Email)
	if err != nil {
		return errors.New("invalid email")
	}

	sevenOrMore, number, upper := verifyPassword(body.Password)

	if !sevenOrMore {
		return errors.New("password must be at least 7 characters")
	}

	if !number {
		return errors.New("password must contain at least one number")
	}

	if !upper {
		return errors.New("password must contain at least one uppercase letter")
	}

	return nil
}

func verifyPassword(s string) (sevenOrMore, number, upper bool) {
	letters := 0
	for _, c := range s {
		switch {
		case unicode.IsNumber(c):
			number = true
			letters++
		case unicode.IsUpper(c):
			upper = true
			letters++
		case unicode.IsLetter(c) || c == ' ':
			letters++
		default:
			return false, false, false
		}
	}
	sevenOrMore = letters >= 7
	return
}
