package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var (
	upgrader = websocket.Upgrader{}
	clients	[]websocket.Conn
)

func main() {
	app := echo.New()
	app.Use(middleware.Recover())

	app.Renderer = &Template{
		Templates: template.Must(template.ParseGlob("templates/*.html")),
	}

	app.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", nil)
	})

	app.GET("/test", func(c echo.Context) error {
		ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
		if err != nil {
			return err
		}
		defer ws.Close()

		clients := append(clients, *ws)
	
		for {
			_, msg, err := ws.ReadMessage()
			if err != nil {
				c.Logger().Error(err)
			}
			fmt.Printf("%s -- message received\n",msg)

			for idx, client := range clients {
				fmt.Println(idx, "<-----")
				err := client.WriteMessage(websocket.TextMessage, msg)
				if err != nil {
					c.Logger().Error(err)
				}
			}
		}
	})

	app.Logger.Fatal(app.Start(":8080"))
}