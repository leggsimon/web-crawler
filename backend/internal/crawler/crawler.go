package crawler

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

type CrawlResult struct {
	Status int
	Html   string
}

func CrawlUrl(url string) (CrawlResult, error) {
	resp, err := http.Get(url)
	if err != nil {
		// handle error
		log.Printf("failed to get url %s: %v", url, err)
		return CrawlResult{}, fmt.Errorf("failed to get url %s: %w", url, err)
	}
	defer resp.Body.Close()

	// if response is not html, error and fail the crawl
	contentType := resp.Header.Get("Content-Type") // maybe this should use http.DetectContentType?
	if !strings.Contains(contentType, "text/html") {
		err := fmt.Errorf("server returned a non html Content-Type header, got: %s", contentType)
		log.Printf("%v", err)
		return CrawlResult{Status: resp.StatusCode}, err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Printf("failed to read response body for url %s: %v", url, err)
		return CrawlResult{}, fmt.Errorf("failed to read response body for url %s: %w", url, err)
	}
	// make a request for url

	// parse the html

	// store each link

	// for each link make a request and store the status code.
	return CrawlResult{
		Status: resp.StatusCode,
		Html:   string(body),
	}, nil
}
