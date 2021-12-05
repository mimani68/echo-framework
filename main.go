package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {

	app := echo.New()

	app.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, map[string]interface{}{
			"success": "true",
			"data":    nil,
		})
	})

	app.Start("0.0.0.0:3000")

}
