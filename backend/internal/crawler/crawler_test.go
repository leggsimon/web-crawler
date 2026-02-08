package crawler

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFetchHTML_Success(t *testing.T) {
	html := "<html><body><a href='/page1'>Link</a></body></html>"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Header().Add("Content-Type", "text/html")
		w.Write([]byte(html))
	}))
	defer server.Close()

	// Test the crawler with the test server URL
	result, err := FetchHTML(server.URL)
	if err != nil {
		t.Errorf("expected no error for successful crawl, got: %v", err)
	}

	if result != html {
		t.Errorf("expected HTML content %s, got %s", html, result)
	}
}

func TestFetchHTML_NonSuccessWithHtml(t *testing.T) {
	html := "<html><body><h1>Not Found</h1></body></html>"
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusNotFound)
		w.Header().Add("Content-Type", "text/html")
		w.Write([]byte(html))
	}))
	defer server.Close()

	// Test the crawler with the test server URL
	result, err := FetchHTML(server.URL)

	if err != nil {
		t.Errorf("expected no error for successful crawl, got: %v", err)
	}

	if result != html {
		t.Errorf("expected HTML content %s, got %s", html, result)
	}
}

func TestFetchHTML_NoContentTYpeHeader(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Ok"))
	}))
	defer server.Close()

	// Test the crawler with the test server URL
	result, err := FetchHTML(server.URL)
	if err == nil {
		t.Error("expected error for server error status, got nil")
	}
	msg := err.Error()
	wantedMsg := "server returned a non html Content-Type header, got: text/plain; charset=utf-8"
	if err != nil && msg != wantedMsg {
		t.Errorf("expected correct error message '%s', got: %s", wantedMsg, msg)
	}
	if result != "" {
		t.Errorf("expected empty result on error, got html %s", result)
	}
}

func TestFetchHTML_InvalidURL(t *testing.T) {
	// Test with an invalid URL
	result, err := FetchHTML("http://invalid-url-that-does-not-exist.local")
	if err == nil {
		t.Error("expected error for invalid URL, got nil")
	}
	msg := err.Error()
	wantedMsg := `failed to get url http://invalid-url-that-does-not-exist.local: Get "http://invalid-url-that-does-not-exist.local": dial tcp: lookup invalid-url-that-does-not-exist.local: no such host`
	if err != nil && msg != wantedMsg {
		t.Errorf("expected correct error message '%s', got: %s", wantedMsg, msg)
	}
	if result != "" {
		t.Errorf("expected empty result on error, got html %s", result)
	}
}
