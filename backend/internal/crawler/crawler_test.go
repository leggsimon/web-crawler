package crawler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCrawlUrl_Success(t *testing.T) {
	html := "<html><body><a href='/page1'>Link</a></body></html>"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "text/html")
		w.Write([]byte(html))
	}))
	defer server.Close()

	// Test the crawler with the test server URL
	result, err := CrawlUrl(server.URL)
	if err != nil {
		t.Errorf("expected no error for successful crawl, got: %v", err)
	}

	if result.Status != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, result.Status)
	}
	if result.Html != html {
		t.Errorf("expected HTML content %s, got %s", html, result.Html)
	}
}

func TestCrawlUrl_NonSuccessWithHtml(t *testing.T) {
	html := "<html><body><h1>Not Found</h1></body></html>"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Add("Content-Type", "text/html")
		w.Write([]byte(html))
	}))
	defer server.Close()

	// Test the crawler with the test server URL
	result, err := CrawlUrl(server.URL)

	if err != nil {
		t.Errorf("expected no error for successful crawl, got: %v", err)
	}

	if result.Status != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, result.Status)
	}
	if result.Html != html {
		t.Errorf("expected HTML content %s, got %s", html, result.Html)
	}
}

func TestCrawlUrl_ServerErrorNoContentTYpeHeader(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
	}))
	defer server.Close()

	// Test the crawler with the test server URL
	result, err := CrawlUrl(server.URL)
	if err == nil {
		t.Error("expected error for server error status, got nil")
	}
	msg := err.Error()
	wantedMsg := "server returned a non html Content-Type header, got: text/plain; charset=utf-8"
	if err != nil && msg != wantedMsg {
		t.Errorf("expected correct error message '%s', got: %s", wantedMsg, msg)
	}
	if result.Status != http.StatusInternalServerError {
		t.Errorf("expected status %d, got %d", http.StatusInternalServerError, result.Status)
	}
}

func TestCrawlUrl_InvalidURL(t *testing.T) {
	// Test with an invalid URL
	result, err := CrawlUrl("http://invalid-url-that-does-not-exist.local")
	if err == nil {
		t.Error("expected error for invalid URL, got nil")
	}
	msg := err.Error()
	wantedMsg := `failed to get url http://invalid-url-that-does-not-exist.local: Get "http://invalid-url-that-does-not-exist.local": dial tcp: lookup invalid-url-that-does-not-exist.local: no such host`
	if err != nil && msg != wantedMsg {
		t.Errorf("expected correct error message '%s', got: %s", wantedMsg, msg)
	}
	if result.Status != 0 {
		t.Errorf("expected empty result on error, got status %d", result.Status)
	}
}
