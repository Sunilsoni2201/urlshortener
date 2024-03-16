package models

import (
	"net/url"
	"sync"

	"github.com/Sunilsoni2201/urlshortener/errors"
)

// data structure to store all the shortened URLs in memory
type MemoryDB struct {
	mutex         sync.RWMutex
	UrlsShortened map[string]string
	HostHitMetric map[string]int64
}

// Inititializing  a new instance of MemoryDB with an empty UrlsShortened and HostHitMetric maps
func NewMemoryDB() UrlShortenedRepository {
	db := &MemoryDB{
		UrlsShortened: make(map[string]string),
		HostHitMetric: make(map[string]int64),
	}
	return db
}

// Gets longUrl from the memory database by its shortUrl
func (db *MemoryDB) Get(shortUrl string) (longUrl string, appArr *errors.AppError) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	ok := false
	if longUrl, ok = db.UrlsShortened[shortUrl]; !ok {
		appArr = errors.NewNotFoundError("shortUrl not found")
	}
	return
}

// updates memory database with a new pair of shortUrl:longUrl and updates the HostHitMetric map for longUrl's host
func (db *MemoryDB) Set(shortUrl string, longUrl string) (appArr *errors.AppError) {
	db.mutex.Lock()
	defer db.mutex.Unlock()

	db.UrlsShortened[shortUrl] = longUrl

	u, err := url.Parse(longUrl)
	if err != nil {
		appArr = errors.NewError(err.Error())
		return
	}
	db.HostHitMetric[u.Host]++

	return
}

// Gets the top N hosts with the number of times they have been hit
func (db *MemoryDB) GetTopMetric(n int) (topN map[string]int64, appArr *errors.AppError) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	return
}

// func (db *MemoryDB) LongUrlExist(longUrl string) (shortUrl string, found bool) {
// 	db.mutex.RLock()
// 	defer db.mutex.RUnlock()

// 	for k, v := range db.UrlsShortened {
// 		if v == longUrl {
// 			shortUrl = k
// 			found = true
// 			return
// 		}
// 	}
// 	return
// }

// func (db *MemoryDB) GetMetric(n int) (metric map[string]int64) {
// 	metric = make(map[string]int64)
// 	return
// }

// func (db *MemoryDB) updateMetric(key string) {
// 	db.Metric[key]++
// }
