package models

import (
	"sync"

	"github.com/Sunilsoni2201/urlshortener/errors"
)

type MemoryDB struct {
	mutex         sync.RWMutex
	ShortenedURLs map[string]string
	// Metric        map[string]int64
}

func NewMemoryDB() UrlShortnerDB {
	db := &MemoryDB{
		ShortenedURLs: make(map[string]string),
		// Metric:        make(map[string]int64),
	}
	return db
}

func (db *MemoryDB) Get(shortUrl string) (longUrl string, err *errors.AppError) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	ok := false
	if longUrl, ok = db.ShortenedURLs[shortUrl]; !ok {
		err = errors.NewNotFoundError("key not found in database")
	}
	return
}

func (db *MemoryDB) Set(shortUrl string, longUrl string) (err *errors.AppError) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.ShortenedURLs[shortUrl] = longUrl
	return
}

func (db *MemoryDB) LongUrlExist(longUrl string) (shortUrl string, found bool) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	for k, v := range db.ShortenedURLs {
		if v == longUrl {
			shortUrl = k
			found = true
			return
		}
	}
	return
}

// func (db *MemoryDB) GetMetric(n int) (metric map[string]int64) {
// 	metric = make(map[string]int64)
// 	return
// }

// func (db *MemoryDB) updateMetric(key string) {
// 	db.Metric[key]++
// }
