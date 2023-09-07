package controller

import (
	"echo/config"
	"echo/model"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func GetStores(c echo.Context) error {
    db := config.GetDB()
    var stores []model.Store
    if err := db.Find(&stores).Error; err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch stores"})
    }
    return c.JSON(http.StatusOK, stores)
}

func GetStoreDetail(c echo.Context) error {
    db := config.GetDB()
    storeID := c.Param("id")

    var store model.Store
    if err := db.First(&store, storeID).Error; err != nil {
        return c.JSON(http.StatusNotFound, map[string]string{"message": "Store not found"})
    }

    weatherData, err := getWeatherData(store.Latitude, store.Longitude)
    if err != nil {
        log.Error("Failed to fetch weather data:", err)
        return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Failed to fetch weather data"})
    }

    store.Weather = weatherData

    return c.JSON(http.StatusOK, store)
}

func getWeatherData(latitude, longitude float64) (string, error) {
    apiKey := "5fd7d34f32mshe3c5d7e59e496edp15c2f3jsn1a5c729c7b29"
    endpoint := fmt.Sprintf("https://rapidapi.com/apininjas/api/weather-by-api-ninjas/?latitude=%f&longitude=%f", latitude, longitude)

    req, err := http.NewRequest("GET", endpoint, nil)
    if err != nil {
        return "", err
    }

    req.Header.Set("X-RapidAPI-Key", apiKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return "", fmt.Errorf("failed to fetch weather data: %s", resp.Status)
    }

    body, err := io.ReadAll(resp.Body)
    if err != nil {
        return "", err
    }

    return string(body), nil
}