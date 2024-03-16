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

func (u *urlShortner) GetActualURL(shortUrl string) (longUrl string, err *errors.AppError) {
	if shortUrl == "" {
		logger.Error("Empty short URL provided")
		err = errors.NewInvalidInputError("Empty short URL")
		return
	}

	longUrl, err = u.db.Get(shortUrl)
	if err != nil {
		logger.Error("Unexpected error from database")
		err = errors.NewError("Unexpected error from database")
	}

	if longUrl == "" {
		logger.Error("Wrong short URL")
		err = errors.NewInvalidInputError("Wrong short URL")
	}
	return
}

func (u *urlShortner) CreateShortURL(longUrl string) (shortUrl string, err *errors.AppError) {
	// Check if the longUrl is empty, return an error
	if longUrl == "" {
		err = errors.NewInvalidInputError("Empty long URL")
		return
	}

	// Check if the longUrl is already in the DB, return the existing shortUrl
	var found bool
	if shortUrl, found = u.db.LongUrlExist(longUrl); found {
		return
	}

	i, retryCount := 0, 100
	for i < retryCount {
		shortUrl = createURLHash(longUrl, 6)
		if _, err1 := u.db.Get(shortUrl); err1 != nil {
			err = u.db.Set(shortUrl, longUrl)
			if err != nil {
				err = errors.NewError("Unexpected error from database while setting the key-value pair.")
			}
		}
	}

	if i == retryCount {
		err = errors.NewError("Max retry limit exceeded to generate short url")
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
