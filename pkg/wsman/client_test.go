package wsman

import (
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"
)

func TestNewClient(t *testing.T) {
	target := "https://example.com/wsman"
	username := "user"
	password := "password"
	useDigest := false

	client := NewClient(target, username, password, useDigest)

	if client.endpoint != target {
		t.Errorf("Expected endpoint to be %s, but got %s", target, client.endpoint)
	}
	if client.username != username {
		t.Errorf("Expected username to be %s, but got %s", username, client.username)
	}
	if client.password != password {
		t.Errorf("Expected password to be %s, but got %s", password, client.password)
	}
	if client.useDigest != useDigest {
		t.Errorf("Expected useDigest to be %v, but got %v", useDigest, client.useDigest)
	}
}

func TestClient_Post(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", ContentType)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("<SampleResponse>OK</SampleResponse>"))
	}))
	defer ts.Close()

	target := ts.URL
	username := "user"
	password := "password"
	useDigest := false

	client := NewClient(target, username, password, useDigest)
	msg := "<SampleRequest>Request</SampleRequest>"

	response, err := client.Post(msg)
	if err != nil {
		t.Errorf("Unexpected error during POST: %v", err)
	}

	expectedResponse := "<SampleResponse>OK</SampleResponse>"
	if string(response) != expectedResponse {
		t.Errorf("Expected response to be %s, but got %s", expectedResponse, response)
	}
}
func newMockDigestAuthHandler(username, password string, handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if strings.HasPrefix(authHeader, "Digest ") {
			// Check for the correct username and password in the Authorization header
			if strings.Contains(authHeader, `username="`+username+`"`) { // &&strings.Contains(authHeader, `uri="`+r.URL.RequestURI()
				handler.ServeHTTP(w, r)
			} else {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
			}
		} else {
			// Simulate a server requesting digest authentication with required fields
			w.Header().Set("WWW-Authenticate", `Digest realm="example.com", nonce="mock-nonce", qop="auth", opaque="opaque-data", algorithm=MD5`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
}
func TestClient_PostWithDigestAuth(t *testing.T) {
	// Use a simple digest auth implementation for testing purposes

	ts := httptest.NewServer(newMockDigestAuthHandler("user", "password", http.Handler(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", ContentType)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("<SampleResponse>OK</SampleResponse>"))
	}))))
	defer ts.Close()

	target := ts.URL
	username := "user"
	password := "password"
	useDigest := true

	client := NewClient(target, username, password, useDigest)
	msg := "<SampleRequest>Request</SampleRequest>"

	response, err := client.Post(msg)
	if err != nil {
		t.Errorf("Unexpected error during POST with digest auth: %v", err)
	}

	expectedResponse := "<SampleResponse>OK</SampleResponse>"
	if string(response) != expectedResponse {
		t.Errorf("Expected response to be %s, but got %s", expectedResponse, response)
	}
}

func TestClient_PostWithDigestAuthUnauthorized(t *testing.T) {
	ts := httptest.NewServer(newMockDigestAuthHandler("user", "password", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", ContentType)
		w.WriteHeader(http.StatusOK)
	})))
	defer ts.Close()

	target := ts.URL
	username := "wronguser"
	password := "wrongpassword"
	useDigest := true

	client := NewClient(target, username, password, useDigest)
	msg := "<SampleRequest>Request</SampleRequest>"

	_, err := client.Post(msg)
	if err == nil {
		t.Error("Expected error during POST with wrong digest auth credentials, but got nil")
	}
}

func TestClient_PostWithBasicAuth(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok || username != "user" || password != "password" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		w.Header().Set("Content-Type", ContentType)
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("<SampleResponse>OK</SampleResponse>"))
	}))
	defer ts.Close()

	target := ts.URL
	username := "user"
	password := "password"
	useDigest := false

	client := NewClient(target, username, password, useDigest)
	msg := "<SampleRequest>Request</SampleRequest>"

	response, err := client.Post(msg)
	if err != nil {
		t.Errorf("Unexpected error during POST with basic auth: %v", err)
	}

	expectedResponse := "<SampleResponse>OK</SampleResponse>"
	if string(response) != expectedResponse {
		t.Errorf("Expected response to be %s, but got %s", expectedResponse, response)
	}
}
func TestClient_PostUnauthorized(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
	}))
	defer ts.Close()

	target := ts.URL
	username := "wronguser"
	password := "wrongpassword"
	useDigest := false

	client := NewClient(target, username, password, useDigest)
	msg := "<SampleRequest>Request</SampleRequest>"

	_, err := client.Post(msg)
	if err == nil {
		t.Error("Expected error during POST with wrong credentials, but got nil")
	}
}

func TestClient_PostInvalidResponse(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", ContentType)
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Internal Server Error"))
	}))
	defer ts.Close()

	target := ts.URL
	username := "user"
	password := "password"
	useDigest := false

	client := NewClient(target, username, password, useDigest)
	msg := "<SampleRequest>Request</SampleRequest>"

	_, err := client.Post(msg)
	if err == nil {
		t.Error("Expected error during POST with invalid response, but got nil")
	}
}
