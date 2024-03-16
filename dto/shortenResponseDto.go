package dto

type ShortenResponse struct {
	ShortUrl string `json:"url"`
}

func NewShortenResponse(u string) *ShortenResponse {
	return &ShortenResponse{ShortUrl: u}
}

func (resp ShortenResponse) GetUrl() string {
	return resp.ShortUrl
}
