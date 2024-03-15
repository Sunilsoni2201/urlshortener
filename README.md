# urlshortener

URL Shortener Service
This project aims to build a simple URL shortener service that accepts a URL as an argument via a REST API and returns a shortened URL as a result. The goal is to create an API-only version of a URL shortener, similar to services like Bitly(https://bitly.com/).

Features
**Custom URL Shortening:** The service will generate a shortened URL for a given input URL. It will not rely on external shortening APIs but instead implement this functionality internally.

**Redirection:** Users clicking on the shortened URL will be redirected to the original URL.

**In Memory Storage:** The application will store both the original and shortened URLs in memory for efficient retrieval.

**Duplicate URL Handling:** If the same URL is submitted multiple times, the service will return the previously generated shortened URL instead of creating a new one.

**Metrics API:** The service will include a metrics API that returns the top 3 domain names that have been shortened the most number of times.

**Usage**
To use the URL shortener service, send a POST request to the /shorten endpoint with the original URL as the payload. The response will contain the shortened URL.

For redirection, users can visit the shortened URL, which will automatically redirect them to the original URL.

To retrieve metrics on the most frequently shortened domains, send a GET request to the /metrics endpoint.