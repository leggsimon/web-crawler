package crawler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func FetchHTML(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		log.Printf("failed to get url %s: %v", url, err)
		return "", fmt.Errorf("failed to get url %s: %w", url, err)
	}
	defer resp.Body.Close()

	// if response is not html, error and fail the crawl
	contentType := resp.Header.Get("Content-Type") // maybe this should use http.DetectContentType?
	if !strings.Contains(contentType, "text/html") {
		err := fmt.Errorf("server returned a non html Content-Type header, got: %s", contentType)
		log.Printf("%v", err)
		return "", err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read response body for url %s: %v", url, err)
		return "", fmt.Errorf("failed to read response body for url %s: %w", url, err)
	}

	return string(body), nil
}

func CrawlUrl(url string) {
	// make a request for url
	// html, err := FetchHTML(url)

	// parse the html

	// store each link

	// for each link make a request and store the status code.
}
