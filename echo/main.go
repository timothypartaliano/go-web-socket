package main

import (
	"echo/config"
	"echo/controller"
	"echo/middleware"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	// e.Use(middleware.Recover())

	customValidator := validator.New()
    e.Validator = &CustomValidator{validator: customValidator}

	e.POST("/users/register", controller.RegisterUser)
	e.POST("/users/login", controller.LoginUser)
	e.GET("/stores", controller.GetStores, middleware.AuthenticateUser)
    e.GET("/stores/:id", controller.GetStoreDetail, middleware.AuthenticateUser)

	_, err := config.InitDB()
	if err != nil {
		e.Logger.Fatal(err)
	}

	e.Start(":8080")
}

type CustomValidator struct {
    validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
    if err := cv.validator.Struct(i); err != nil {
        return echo.NewHTTPError(http.StatusBadRequest, err.Error())
    }
    return nil
}