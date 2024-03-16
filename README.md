# URL Shortener Service
This project aims to build a simple URL shortener service that accepts a URL as an argument via a REST API and returns a shortened URL as a result. The goal is to create an API-only version of a URL shortener, similar to services like Bitly(https://bitly.com/).

# Features
**Custom URL Shortening:** The service will generate a shortened URL for a given input URL. It will not rely on any external shortening APIs but the functionality is implemented internally.

**Redirection:** Users clicking on the shortened URL will be redirected to the original URL.

**Duplicate URL Handling:** If the same URL is submitted multiple times, the service will return the previously generated shortened URL instead of creating a new one.

**Metrics API:** The service will include a metrics API that returns the top 3 domain names that have been shortened the most number of times.

**In Memory Storage:** The application will store both the original and shortened URLs in memory for efficient retrieval and once service is stopped data is lost.
Extension: we  can add a feature that allows users to persist the data in DB or filesystem.

# Usage

**Shorten URL**
Endpoint: POST /shorten

This endpoint shortens a long URL provided in the request body and returns the shortened URL.

Request -->

`
curl -X POST <machine_outbound_ip>:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://www.google.com/"}'
`

Response -->

`
{
  "url": "<machine_outbound_ip>:8080/c03aa1"
}
`

**Redirect to Original URL**

Endpoint: GET /:shortUrl

This endpoint redirects the user to the original URL associated with the provided shortUrl.

Request -->

`
curl -X GET <machine_outbound_ip>:8080/c03aa1
`

Response -->
redirect user to original long url, in this case "https://www.google.com/"


**Top Metrics**
Endpoint: GET /topmetric

This endpoint returns the top 3 websites whose URLs have been most frequently shortened using this service.

Request -->

`
curl -X GET http://localhost:8080/topmetric
`

Response -->

`
{"top_metrics":{"www.facebook.com":3,"www.google.com":2, "www.youtube.com":3}}
`

# Setup
To run this service locally:

Clone this repository.
Build the executable: go build -o bin/url-shortener main.go
Run the executable: ./url-shortener
Make sure you have Go installed on your system.
