package controller

import (
	"echo/config"
	"echo/model"
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterUser(c echo.Context) error {
    user := new(model.User)
    if err := c.Bind(user); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid request body"})
    }

    if err := c.Validate(user); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{"message": err.Error()})
    }

    db := config.GetDB()

    if err := db.Create(user).Error; err != nil {
        fmt.Println("Error inserting user:", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to insert user"})
    }

    return c.JSON(http.StatusOK, map[string]string{"message": "Registration successful"})
}