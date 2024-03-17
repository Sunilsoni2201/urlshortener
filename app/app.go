package app

import (
	"github.com/Sunilsoni2201/urlshortener/handlers"
	"github.com/Sunilsoni2201/urlshortener/models"
	"github.com/Sunilsoni2201/urlshortener/services"
	"github.com/labstack/echo"
)

func Start() {
	// Echo instance
	e := echo.New()

	// Middleware
	// e.Use(middleware.Logger())
	// e.Use(middleware.Recover())

	db := models.NewMemoryDB()
	srv := services.NewUrlShortenerService(db)
	handler := handlers.NewUrlshortenerHandler(srv)

	// API to create a new short URL from long URL
	e.POST("/shorten", handler.Shorten)

	// API to redirect user to the original URL if it exists
	e.GET("/:shortUrl", handler.RedirectToOriginalURL)

	// API to get top metric (currently it is hardcoded to 3)
	e.GET("/topmetric", handler.GetTopMetric)

	e.Logger.Fatal(e.Start(":8080"))
}
