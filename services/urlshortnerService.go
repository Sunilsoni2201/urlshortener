package services

import (
	"crypto/sha256"
	"fmt"
	"time"

	"github.com/Sunilsoni2201/urlshortener/errors"
	"github.com/Sunilsoni2201/urlshortener/logger"
	"github.com/Sunilsoni2201/urlshortener/models"
)

type UrlShortner interface {
	GetActualURL(string) (string, *errors.AppError)
	CreateShortURL(string) (string, *errors.AppError)
}

type urlShortner struct {
	db models.UrlShortenedRepository
}

func NewUrlShortenerService(db models.UrlShortenedRepository) *urlShortner {
	return &urlShortner{
		db: db,
	}
}

func (u *urlShortner) GetActualURL(shortUrl string) (longUrl string, appErr *errors.AppError) {
	if shortUrl == "" {
		logger.Error("Empty shortUrl provided")
		appErr = errors.NewInvalidInputError("Empty shortUrl")
		return longUrl, appErr
	}

	longUrl, appErr = u.db.Get(shortUrl)
	if appErr != nil {
		logger.Error("Get Error: " + appErr.Message)
		appErr = errors.NewError("Unexpected error from database")
		return longUrl, appErr
	}

	if longUrl == "" {
		logger.Error("longUrl not found for given shortUrl: " + shortUrl)
		appErr = errors.NewInvalidInputError("Provided shortUrl not found in the system")
		return longUrl, appErr
	}
	return
}

func (u *urlShortner) CreateShortURL(longUrl string) (shortUrl string, appErr *errors.AppError) {
	// Check if the longUrl is empty, return an error
	if longUrl == "" {
		appErr = errors.NewInvalidInputError("Empty longUrl provided")
		return "", appErr
	}

	// Check if the longUrl is already in the DB, return the existing shortUrl
	var found bool
	if shortUrl, found = u.db.LongUrlExist(longUrl); found {
		logger.Info("longUrl already exists in  the system with shortUrl : " + shortUrl)
		return shortUrl, nil
	}

	// There are chance that
	var i int
	retryCount := 100
	for i = 0; i < retryCount; i++ {
		shortUrl = createURLHash(longUrl, 6)
		_, err1 := u.db.Get(shortUrl)
		if err1 != nil {
			logger.Info("longUrl: " + longUrl + ", shortUrl: " + shortUrl)
			appErr = u.db.Set(shortUrl, longUrl)
			if appErr != nil {
				appErr = errors.NewError("Unexpected error from database while setting the key-value pair.")
			}
			break
		}
	}

	if i == retryCount {
		appErr = errors.NewError("Max retry limit exceeded to generate short url")
	}
	return
}

func createURLHash(longUrl string, hashLen int) (shortUrl string) {
	if hashLen == 0 {
		return ""
	}

	salt := time.Now().String()
	sha := sha256.Sum256([]byte(longUrl + salt))
	return fmt.Sprintf("%x", sha)[:hashLen]
}
