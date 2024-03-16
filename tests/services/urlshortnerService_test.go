package tests

import (
	"testing"

	"github.com/Sunilsoni2201/urlshortener/models"
	"github.com/Sunilsoni2201/urlshortener/services"
)

func TestNewUrlShortenerService(t *testing.T) {
	service := services.NewUrlShortenerService(models.NewMemoryDB())

}
