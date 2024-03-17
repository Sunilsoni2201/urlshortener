package models

import (
	"net/url"
	"sort"
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
		appArr = errors.NewNotFoundError("shortUrl not found: " + shortUrl)
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

	if len(db.HostHitMetric) < n {
		return db.HostHitMetric, nil
	}

	topN = sortHostHitMetric(db.HostHitMetric, n)
	return topN, nil
}

// Check if provided longUrl exists in the system
func (db *MemoryDB) LongUrlExist(longUrl string) (shortUrl string, found bool) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()

	for k, v := range db.UrlsShortened {
		if v == longUrl {
			shortUrl = k
			found = true
			return
		}
	}
	return
}

// Pair represents a key-value pair
type Pair struct {
	Key   string
	Value int64
}

// Sorts the host hit metric by value in descending order and returns the sorted list as a map of size n
func sortHostHitMetric(hostHitMetric map[string]int64, n int) (sortedHostHitMetric map[string]int64) {
	// Convert map to slice of key-value pairs
	var pairs []Pair
	for k, v := range hostHitMetric {
		pairs = append(pairs, Pair{k, v})
	}

	// Sort the slice by values
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].Value > pairs[j].Value
	})

	// Create a new map from the sorted slice
	sortedHostHitMetric = make(map[string]int64)
	count := 0
	for _, pair := range pairs {
		sortedHostHitMetric[pair.Key] = pair.Value
		count++
		if count == n {
			break
		}
	}
	return sortedHostHitMetric
}
