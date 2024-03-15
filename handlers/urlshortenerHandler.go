package handlers

import (
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

	surl := c.Param("shortUrl")
	ourl, appErr := h.service.GetActualURL(surl)
	if appErr != nil {
		logger.Error(appErr.Message)
		return c.String(http.StatusNotFound, appErr.Message)
	}

	c.Response().Header().Set("Location", ourl)

	return c.String(http.StatusMovedPermanently, "")
}
