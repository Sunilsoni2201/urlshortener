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

	e.POST("/shorten", handler.Shorten)
	e.GET("/:shortUrl", handler.GetLongUrl)
	e.GET("/topmetric", handler.GetTopMetric)

	e.Logger.Fatal(e.Start(":8080"))
}
