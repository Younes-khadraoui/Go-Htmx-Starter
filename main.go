package main

import (
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"starter/assets"
	"starter/assets/views"
)

func main() {
	server := echo.New()
	server.Use(middleware.Logger())

	// Serve static files
	server.GET("/static/*", func(c echo.Context) error {
		path := c.Param("*")
		data, err := assets.StaticFS.ReadFile("static/" + path)
		if err != nil {
			return echo.NewHTTPError(404, "File not found")
		}
		contentType := http.DetectContentType(data)
		return c.Blob(http.StatusOK, contentType, data)
	})

	// Serve homepage using Templ
	server.GET("/", func(c echo.Context) error {
		return views.HomePage("Gopher").Render(c.Request().Context(), c.Response())
	})

	// HTMX counter endpoint
	var count int
	server.GET("/count", func(c echo.Context) error {
		count++
		html := `<p id="htmx-response">HTMX Count: ` + strconv.Itoa(count) + `</p>`
		return c.HTML(http.StatusOK, html)
	})

	log.Println("Server started at http://localhost:8000")
	if err := server.Start(":8000"); err != nil {
		log.Fatal(err)
	}
}
