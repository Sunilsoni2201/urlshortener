package tests

import (
	"testing"

	"github.com/Sunilsoni2201/urlshortener/models"
)

func TestNewMemoryDB(t *testing.T) {
	db := models.NewMemoryDB()
	shortUrl := "abc123"
	longUrl, err := db.Get(shortUrl)

	if longUrl != "" {
		t.Errorf(`Long url must be empty (""), got "%s"`, longUrl)
	}

	if err == nil {
		t.Errorf("Error must not be nil")
	}
}

func TestGetAndSet(t *testing.T) {
	testcases := []struct {
		name     string
		shortUrl string
		longUrl  string

		getkey    string
		want      string
		wantError bool
	}{
		{
			name:     "valid input",
			shortUrl: "abc123",
			longUrl:  "https://www.google.com",

			getkey:    "abc123",
			want:      "https://www.google.com",
			wantError: false,
		},
		{
			name:     "invalid input",
			shortUrl: "abc123",
			longUrl:  "https://www.google.com",

			getkey:    "abc12",
			want:      "",
			wantError: true,
		},
	}

	db := models.NewMemoryDB()
	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			err := db.Set(tc.shortUrl, tc.longUrl)
			if err != nil {
				t.Errorf("error in set shortUrl : %v, longUrl : %v", tc.shortUrl, tc.longUrl)
			}
			got, err := db.Get(tc.getkey)
			gotError := (err != nil)

			if tc.wantError != gotError {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}

			if tc.want != got {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}

}
