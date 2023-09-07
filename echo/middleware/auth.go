package middleware

import (
    "echo/config"
    "echo/model"
    "net/http"
    "github.com/labstack/echo/v4"
)

func AuthenticateUser(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        username := c.Get("username").(string)

        db := config.GetDB()
        var user model.User
        if err := db.Where("username = ?", username).First(&user).Error; err != nil {
            return echo.NewHTTPError(http.StatusUnauthorized, "Authentication failed")
        }

        c.Set("user", user)

        return next(c)
    }
}
