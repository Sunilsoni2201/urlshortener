package dto

import "net/url"

type ShortenRequest struct {
	LongUrl string `json:"url"`
}

func NewShortenRequest(u string) *ShortenRequest {
	return &ShortenRequest{
		LongUrl: u,
	}
}
func (r *ShortenRequest) GetUrl() string {
	return r.LongUrl
}

func (r *ShortenRequest) IsValidURL() (bool, error) {
	parsedUrl, err := url.Parse(r.LongUrl)
	if err != nil {
		return false, err
	}

	if parsedUrl.Scheme == "" || parsedUrl.Host == "" {
		return false, nil
	}
	return true, nil
}
