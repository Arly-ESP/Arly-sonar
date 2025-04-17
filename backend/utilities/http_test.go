package utilities

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
)

func TestHTTPRequest(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/api/test-route" {
			t.Errorf("Expected URL path /api/test-route, got %s", r.URL.Path)
		}
		if r.Method != "GET" {
			t.Errorf("Expected method GET, got %s", r.Method)
		}
		if r.Header.Get("Content-Type") != "application/json" {
			t.Errorf("Expected header Content-Type: application/json, got %s", r.Header.Get("Content-Type"))
		}
		bodyBytes, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Errorf("Error reading request body: %v", err)
		}
		if string(bodyBytes) != "test body" {
			t.Errorf("Expected request body 'test body', got '%s'", string(bodyBytes))
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Hello World"))
	}))
	defer ts.Close()

	u, err := url.Parse(ts.URL)
	if err != nil {
		t.Fatalf("Error parsing test server URL: %v", err)
	}
	port := u.Port()
	if port == "" {
		t.Fatal("Test server port is empty")
	}

	os.Setenv("PORT", port)

	headers := map[string]string{
		"Content-Type": "application/json",
	}

	resp, err := HTTPRequest("GET", "/test-route", []byte("test body"), headers)
	if err != nil {
		t.Fatalf("HTTPRequest returned error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}

	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Error reading response body: %v", err)
	}
	if strings.TrimSpace(string(respBody)) != "Hello World" {
		t.Errorf("Expected response body 'Hello World', got '%s'", string(respBody))
	}
}
