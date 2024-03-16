package models

import "github.com/Sunilsoni2201/urlshortener/errors"

type UrlShortened struct {
	LongURL  string
	ShortURL string
	Hostname string
}

type UrlShortenedRepository interface {
	Get(string) (string, *errors.AppError)
	Set(string, string) *errors.AppError
	GetTopMetric(int) (map[string]int64, *errors.AppError)
	// updateMetric(string)
	// LongUrlExist(string) (string, bool)
}
