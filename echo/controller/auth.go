package controller

import (
	"echo/config"
	"echo/model"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtSecret = []byte("PARTALIANO")

func LoginUser(c echo.Context) error {
    username := c.FormValue("username")
    password := c.FormValue("password")

    if !authenticateUser(username, password) {
        return echo.NewHTTPError(http.StatusUnauthorized, "Invalid credentials")
    }

    token, err := generateJWTToken(username)
    if err != nil {
        return echo.NewHTTPError(http.StatusInternalServerError, "Failed to generate token")
    }

    return c.JSON(http.StatusOK, map[string]string{
        "token": token,
    })
}

func Authenticate(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        username := c.Get("username").(string)

        db := config.GetDB()
        var user model.User
        if err := db.Where("username = ?", username).First(&user).Error; err != nil {
            if err == gorm.ErrRecordNotFound {
                return echo.NewHTTPError(http.StatusUnauthorized, "User not found")
            }
            return echo.NewHTTPError(http.StatusInternalServerError, "Database error")
        }

        c.Set("user", user)

        return next(c)
    }
}

func generateJWTToken(username string) (string, error) {
    token := jwt.New(jwt.SigningMethodHS256)

    claims := token.Claims.(jwt.MapClaims)
    claims["username"] = username
    claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }

    return tokenString, nil
}

func authenticateUser(username, password string) bool {
    db := config.GetDB()

    var user model.User
    if err := db.Where("username = ?", username).First(&user).Error; err != nil {
        return false
    }

    err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
    if err != nil {
        return false
    }

    return true
}