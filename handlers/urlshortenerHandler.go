package handlers

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

	"github.com/Sunilsoni2201/urlshortener/dto"
	"github.com/Sunilsoni2201/urlshortener/logger"
	"github.com/Sunilsoni2201/urlshortener/services"
	"github.com/Sunilsoni2201/urlshortener/utils"
	"github.com/labstack/echo"
)

type UrlshortenerHandler struct {
	service services.UrlShortner
}

func NewUrlshortenerHandler(srv services.UrlShortner) *UrlshortenerHandler {
	return &UrlshortenerHandler{service: srv}
}

func (h *UrlshortenerHandler) Shorten(c echo.Context) error {
	req := dto.ShortenRequest{}
	_ = c.Bind(&req)

	resp, appErr := h.service.CreateShortURL(&req)
	if appErr != nil {
		logger.Error(appErr.Message)
		return c.String(http.StatusInternalServerError, appErr.Message)
	}

	serverPort := c.Echo().Listener.Addr().(*net.TCPAddr).Port
	shortURL := fmt.Sprintf("%v:%v/%s", utils.GetOutboundIP(), serverPort, resp.GetUrl())
	return c.String(http.StatusOK, shortURL)
}

func (h *UrlshortenerHandler) RedirectToOriginalURL(c echo.Context) error {

	shortUrl := c.Param("shortUrl")
	longUrl, appErr := h.service.GetOrignialLongURL(shortUrl)
	logger.Info("shortUrl: " + shortUrl + ", longUrl: " + longUrl)
	if appErr != nil {
		logger.Error(appErr.Message)
		return c.String(http.StatusNotFound, appErr.Message)
	}
	c.Response().Header().Set("Location", longUrl)
	return c.String(http.StatusMovedPermanently, "")
}

func (h *UrlshortenerHandler) GetTopMetric(c echo.Context) error {
	Top3 := 3
	ans := h.service.GetTopMetric(Top3)

	jsonString, err := json.Marshal(ans)
	if err != nil {
		fmt.Println("Error:", err)
		return nil
	}

	return c.String(http.StatusOK, string(jsonString))
}
