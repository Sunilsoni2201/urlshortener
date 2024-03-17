package services

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/Sunilsoni2201/urlshortener/dto"
	"github.com/Sunilsoni2201/urlshortener/errors"
	"github.com/Sunilsoni2201/urlshortener/logger"
	"github.com/Sunilsoni2201/urlshortener/models"
)

/*
This is service layer  for URL shortening functionality.
It interacts with the database layer and returns appropriate responses to the controller layer(handler fucntions).
*/
type UrlShortner interface {
	GetOrignialLongURL(string) (string, *errors.AppError)
	CreateShortURL(*dto.ShortenRequest) (*dto.ShortenResponse, *errors.AppError)
	GetTopMetric(int) *dto.TopMetricResponse
}

type urlShortner struct {
	db models.UrlShortenedRepository
}

func NewUrlShortenerService(db models.UrlShortenedRepository) *urlShortner {
	return &urlShortner{
		db: db,
	}
}

// Get the original long URL from database using the Short URL.
// If it does not exist in DB then return an error
func (u *urlShortner) GetOrignialLongURL(shortUrl string) (longUrl string, appErr *errors.AppError) {

	longUrl, appErr = u.db.Get(shortUrl)
	if appErr != nil {
		logger.Error("Get Error: " + appErr.Message)
		appErr = errors.NewError("Unexpected error from database")
		return "", appErr
	}

	if longUrl == "" {
		logger.Error("longUrl not found for given shortUrl: " + shortUrl)
		appErr = errors.NewInvalidInputError("Provided shortUrl not found in the system")
		return "", appErr
	}
	return longUrl, nil
}

// Create a new entry in the database with Long URL and generate a unique Short URL.
// Return the Short URL to the client so that they can share this link.
func (u *urlShortner) CreateShortURL(req *dto.ShortenRequest) (resp *dto.ShortenResponse, appErr *errors.AppError) {

	valid, err := req.IsValidURL()
	if err != nil || !valid {
		logger.Error("Error while validating request url :" + req.GetUrl())
		return nil, errors.NewInvalidInputError("Please provide a full valid URL with host & scheme(http or https): " + req.GetUrl())
	}

	var shortUrl string
	var found bool
	// Check if the longUrl is already in the DB, return the existing shortUrl
	if shortUrl, found = u.db.LongUrlExist(req.GetUrl()); found {
		logger.Info("longUrl already exists in  the system with shortUrl : " + shortUrl)
		return dto.NewShortenResponse(shortUrl), nil
	}

	// There are chance that
	var i int
	retryCount := 100
	longUrl := req.GetUrl()
	for i = 0; i < retryCount; i++ {
		shortUrl = createURLHash(longUrl, 6)
		_, err1 := u.db.Get(shortUrl)
		if err1 != nil {
			logger.Info("longUrl: " + longUrl + ", shortUrl: " + shortUrl)
			appErr = u.db.Set(shortUrl, longUrl)
			if appErr != nil {
				appErr = errors.NewError("Unexpected error from database while setting the key-value pair.")
				return nil, appErr
			}
			break
		}
	}

	if i == retryCount {
		appErr = errors.NewError("Max retry limit exceeded to generate short url")
		return nil, appErr
	}

	if shortUrl == "" {
		appErr = errors.NewError("Something went wrong and shortUrl could not be generated for this URL.")
		return nil, appErr
	}

	return dto.NewShortenResponse(shortUrl), nil
}

// createURLHash generate and returns the longUrl hash with given length as shortUrl
func createURLHash(longUrl string, hashLen int) (shortUrl string) {
	if hashLen == 0 {
		return ""
	}

	salt := time.Now().String()
	sha := sha256.Sum256([]byte(longUrl + salt))
	return fmt.Sprintf("%x", sha)[:hashLen]
}

// GetTopMetric gets top N(count) most visited URLs in descending order of visit count
func (u *urlShortner) GetTopMetric(count int) *dto.TopMetricResponse {
	metricsMap, _ := u.db.GetTopMetric(count)
	return dto.NewTopMetricRespons(metricsMap)
}
