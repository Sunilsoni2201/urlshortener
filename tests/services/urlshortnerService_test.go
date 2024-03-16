package tests

import (
	"strconv"
	"testing"

	"github.com/Sunilsoni2201/urlshortener/dto"
	"github.com/Sunilsoni2201/urlshortener/models"
	"github.com/Sunilsoni2201/urlshortener/services"
)

// func TestNewUrlShortenerService(t *testing.T) {
// 	service := services.NewUrlShortenerService(models.NewMemoryDB())

// 	longUrls := []string{
// 		"https://www.facebook.com/",
// 		"https://www.youtube.com/",
// 		"https://www.linkedin.com/",
// 		"https://stackoverflow.com/",
// 		"https://www.instagram.com/",
// 		"https://www.facebook.com/login.php",
// 		"https://www.youtube.com/watch?v=dQw4w9WgXcQ",
// 		"https://www.linkedin.com/login",
// 		"https://www.facebook.com/pages/create/",
// 		"https://www.youtube.com/channel/UCBR8-60-B28hp2BmDPdntcQ",
// 	}

// 	var sortUrls []string
// 	for _, url := range longUrls {
// 		resp, _ := service.CreateShortURL(dto.NewShortenRequest(url))
// 		sortUrls = append(sortUrls, resp.GetUrl())
// 	}

// 	type testcase struct {
// 		name     string
// 		shortUrl string
// 		longUrl  string

// 		getkey    string
// 		want      string
// 		wantError bool
// 	}
// 	testcases := []testcase{}

// 	for i, longUrl := range longUrls {
// 		testcases = append(testcases, testcase{})
// 	}
// }

func BenchmarkCreateShortURL(b *testing.B) {
	shortenerService := services.NewUrlShortenerService(models.NewMemoryDB())

	for i := 0; i < b.N; i++ {
		u := "https://www.google.com/search?q=linkedin" + strconv.Itoa(i)
		_, _ = shortenerService.CreateShortURL(dto.NewShortenRequest(u))
	}
}
