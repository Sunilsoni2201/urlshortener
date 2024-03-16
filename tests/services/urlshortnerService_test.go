package tests

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"testing"

	"github.com/Sunilsoni2201/urlshortener/dto"
	"github.com/Sunilsoni2201/urlshortener/models"
	"github.com/Sunilsoni2201/urlshortener/services"
)

var shortUrlService services.UrlShortner

// setup function to initialize common resource
func setup() {
	shortUrlService = services.NewUrlShortenerService(models.NewMemoryDB())
	fmt.Println("Setup complete")
}

// cleanup function to clean up resources after tests
func cleanup() {
	fmt.Println("Cleanup complete")
}

// TestMain is the entry point for test execution
func TestMain(m *testing.M) {
	setup()         // Call setup function before running tests
	code := m.Run() // Run tests
	cleanup()       // Call cleanup function after running tests
	os.Exit(code)
}

func TestCreateShortUrlWithSingleUrls(t *testing.T) {
	longUrl := "https://www.facebook.com/"
	resp, _ := shortUrlService.CreateShortURL(dto.NewShortenRequest(longUrl))
	gotShortUrl := resp.GetUrl()
	gotLongUrl, _ := shortUrlService.GetOrignialLongURL(gotShortUrl)

	if longUrl != gotLongUrl {
		t.Errorf("Want: %v, Got: %v", longUrl, gotLongUrl)
	}
}

func TestCreateShortUrlWithDifferentUrls(t *testing.T) {
	longUrls := []string{
		"https://www.facebook.com/",
		"https://www.youtube.com/",
		"https://www.linkedin.com/",
		"https://stackoverflow.com/",
		"https://www.instagram.com/",
		"https://www.facebook.com/login.php",
		"https://www.youtube.com/watch?v=dQw4w9WgXcQ",
		"https://www.linkedin.com/login",
		"https://www.facebook.com/pages/create/",
		"https://www.youtube.com/channel/UCBR8-60-B28hp2BmDPdntcQ",
	}

	gotSortUrls := []string{}
	gotLongUrls := []string{}
	for _, url := range longUrls {
		resp, _ := shortUrlService.CreateShortURL(dto.NewShortenRequest(url))
		gotSortUrls = append(gotSortUrls, resp.GetUrl())

		gotLongUrl, _ := shortUrlService.GetOrignialLongURL(resp.GetUrl())
		gotLongUrls = append(gotLongUrls, gotLongUrl)
	}

	for i, _ := range longUrls {
		if longUrls[i] != gotLongUrls[i] {
			t.Errorf("Want: %v, Got: %v", longUrls[i], gotLongUrls[i])
		}
	}

}

func TestCreateShortUrlWithDuplicateUrls(t *testing.T) {
	longUrl := "https://www.facebook.com/"
	resp, _ := shortUrlService.CreateShortURL(dto.NewShortenRequest(longUrl))
	gotShortUrl := resp.GetUrl()

	for i := 0; i < 5; i++ {
		newResp, _ := shortUrlService.CreateShortURL(dto.NewShortenRequest(longUrl))
		newShortUrl := newResp.GetUrl()
		if gotShortUrl != newShortUrl {
			t.Errorf("Want: %v, Got: %v", gotShortUrl, newShortUrl)
		}
	}
}

func TestCreateShortUrlWithInvalideUrls(t *testing.T) {
	longUrl := "httpswww.facebook.com/"
	resp, err := shortUrlService.CreateShortURL(dto.NewShortenRequest(longUrl))
	if resp != nil || err == nil {
		t.Errorf("Expecting error but got response: %+v", resp)
	}
}

func TestCreateShortUrlWithBlankUrls(t *testing.T) {
	longUrl := ""
	resp, err := shortUrlService.CreateShortURL(dto.NewShortenRequest(longUrl))
	if resp != nil || err == nil {
		t.Errorf("Expecting error but got response: %+v", resp)
	}
}

func TestGetTopMetrics(t *testing.T) {
	longUrls := []string{
		"https://www.facebook.com/",
		"https://www.youtube.com/",
		"https://www.linkedin.com/",
		"https://stackoverflow.com/",
		"https://www.instagram.com/",
		"https://www.facebook.com/login.php",
		"https://www.youtube.com/watch?v=dQw4w9WgXcQ",
		"https://www.linkedin.com/login/",
		"https://www.facebook.com/pages/create/",
		"https://www.youtube.com/channel/UCBR8-60-B28hp2BmDPdntcQ",
	}

	gotSortUrls := []string{}
	for _, url := range longUrls {
		resp, _ := shortUrlService.CreateShortURL(dto.NewShortenRequest(url))
		gotSortUrls = append(gotSortUrls, resp.GetUrl())
	}

	gotTopMetricResponse := shortUrlService.GetTopMetric(3)
	wantTopMetric := map[string]int64{
		"www.facebook.com": 3,
		"www.youtube.com":  3,
		"www.linkedin.com": 2,
	}

	if !reflect.DeepEqual(wantTopMetric, gotTopMetricResponse.TopMetrics) {
		t.Errorf("Want: %v, Got: %v", wantTopMetric, gotTopMetricResponse.TopMetrics)
	}
}

func BenchmarkCreateShortURL(b *testing.B) {
	shortenerService := services.NewUrlShortenerService(models.NewMemoryDB())

	for i := 0; i < b.N; i++ {
		u := "https://www.google.com/search?q=linkedin" + strconv.Itoa(i)
		_, _ = shortenerService.CreateShortURL(dto.NewShortenRequest(u))
	}
}
