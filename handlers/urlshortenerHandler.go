package handlers

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"

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
	type In struct {
		Url string `json:"url"`
	}
	in := In{}

	_ = c.Bind(&in)

	logger.Info("shortUrl fetched from request(echo.Context) :" + in.Url)

	valid, err := utils.IsValidURL(in.Url)
	if err != nil || !valid {
		return c.String(http.StatusUnprocessableEntity, "Please provide a full valid URL with host & scheme(http or https)")
	}

	surl, appErr := h.service.CreateShortURL(in.Url)
	if appErr != nil {
		logger.Error(appErr.Message)
		return c.String(http.StatusInternalServerError, appErr.Message)
	}

	serverPort := c.Echo().Listener.Addr().(*net.TCPAddr).Port
	shortURL := fmt.Sprintf("%v:%v/%s", utils.GetOutboundIP(), serverPort, surl)
	return c.String(http.StatusOK, shortURL)
}

func (h *UrlshortenerHandler) GetLongUrl(c echo.Context) error {

	shortUrl := c.Param("shortUrl")
	longUrl, appErr := h.service.GetActualURL(shortUrl)
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
