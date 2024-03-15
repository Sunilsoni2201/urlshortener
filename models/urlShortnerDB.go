package models

import "github.com/Sunilsoni2201/urlshortener/errors"

type UrlShortnerDB interface {
	Get(string) (string, *errors.AppError)
	Set(string, string) *errors.AppError
	LongUrlExist(string) (string, bool)
	// updateMetric(string)
	// GetMetric(int) map[string]int64
}
