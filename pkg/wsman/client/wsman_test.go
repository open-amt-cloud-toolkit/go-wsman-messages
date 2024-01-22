/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package client

import (
	"strings"
	"testing"

	"net/http"
	"net/http/httptest"
)

func TestNewClient(t *testing.T) {
	target := "example.com"
	expectedTarget := "http://example.com:16992/wsman"
	username := "user"
	password := "password"
	useDigest := false
	useTLS := false
	selfSignedAllowed := false

	client := NewWsman(target, username, password, useDigest, useTLS, selfSignedAllowed)

	if client.endpoint != expectedTarget {
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

func TestNewClient_TLS(t *testing.T) {
	target := "example.com"
	expectedTarget := "https://example.com:16993/wsman"
	username := "user"
	password := "password"
	useDigest := false
	useTLS := true
	selfSignedAllowed := true

	client := NewWsman(target, username, password, useDigest, useTLS, selfSignedAllowed)

	if client.endpoint != expectedTarget {
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
	useTLS := false
	selfSignedAllowed := false

	client := NewWsman(target, username, password, useDigest, useTLS, selfSignedAllowed)
	msg := "<SampleRequest>Request</SampleRequest>"

	client.endpoint = ts.URL

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
	useTLS := false
	selfSignedAllowed := false

	client := NewWsman(target, username, password, useDigest, useTLS, selfSignedAllowed)
	msg := "<SampleRequest>Request</SampleRequest>"

	client.endpoint = ts.URL
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
	useTLS := false
	selfSignedAllowed := false

	client := NewWsman(target, username, password, useDigest, useTLS, selfSignedAllowed)
	msg := "<SampleRequest>Request</SampleRequest>"

	client.endpoint = ts.URL
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
	useTLS := false
	selfSignedAllowed := false

	client := NewWsman(target, username, password, useDigest, useTLS, selfSignedAllowed)
	msg := "<SampleRequest>Request</SampleRequest>"

	client.endpoint = ts.URL
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
	useTLS := false
	selfSignedAllowed := false

	client := NewWsman(target, username, password, useDigest, useTLS, selfSignedAllowed)
	msg := "<SampleRequest>Request</SampleRequest>"

	client.endpoint = ts.URL
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
	useTLS := false
	selfSignedAllowed := false

	client := NewWsman(target, username, password, useDigest, useTLS, selfSignedAllowed)
	msg := "<SampleRequest>Request</SampleRequest>"

	client.endpoint = ts.URL
	_, err := client.Post(msg)
	if err == nil {
		t.Error("Expected error during POST with invalid response, but got nil")
	}
}

func TestClient_PostWithDigestBlankRealm(t *testing.T) {
	ts := httptest.NewServer(newMockDigestAuthHandler("user", "password", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if strings.HasPrefix(authHeader, "Digest ") {
			//Simulate internal server error
			w.WriteHeader(http.StatusInternalServerError)
		} else {
			// Simulate a server requesting digest authentication with required fields
			w.Header().Set("WWW-Authenticate", `Digest realm="example.com", nonce="mock-nonce", qop="auth", opaque="opaque-data", algorithm=MD5`)
			w.WriteHeader(http.StatusUnauthorized)
		}
	})))
	defer ts.Close()

	target := ts.URL
	username := "user"
	password := "password"
	useDigest := true
	useTLS := false
	selfSignedAllowed := false

	client := NewWsman(target, username, password, useDigest, useTLS, selfSignedAllowed)
	client.challenge.Realm = ""
	msg := "<SampleRequest>Request</SampleRequest>"

	client.endpoint = ts.URL
	_, err := client.Post(msg)
	if err == nil {
		t.Error("Expected error during POST with wrong digest auth credentials, but got nil")
	}
	if !strings.Contains(err.Error(), "500 Internal Server Error") {
		t.Error("Wsman client should not send digest on initial challenges")
	}

}

func TestClient_ProxyUrlTransport(t *testing.T) {
	target := "example.com"
	username := "user"
	password := "password"
	useDigest := true
	useTLS := false
	selfSignedAllowed := false

	client := NewWsman(target, username, password, useDigest, useTLS, selfSignedAllowed)
	err := client.ProxyUrl("http://localhost:3128")
	if err != nil {
		t.Error("Failed to set proxy on proper Transport")
	}
}

func TestClient_InvalidProxyUrlGoodTransport(t *testing.T) {
	target := "example.com"
	username := "user"
	password := "password"
	useDigest := true
	useTLS := false
	selfSignedAllowed := false

	client := NewWsman(target, username, password, useDigest, useTLS, selfSignedAllowed)
	err := client.ProxyUrl("localhost")
	if err == nil {
		t.Error("Failed to detect invalid proxy url")
	}
}

// inline struct for mock roundtripper
type rt struct{}

func (*rt) RoundTrip(r *http.Request) (*http.Response, error) {
	return nil, nil
}

func TestClient_SimpleRountripper(t *testing.T) {
	target := "example.com"
	username := "user"
	password := "password"
	useDigest := true
	useTLS := false
	selfSignedAllowed := false
	mockrt := rt{}
	client := NewWsman(target, username, password, useDigest, useTLS, selfSignedAllowed)
	client.Transport = &mockrt
	err := client.ProxyUrl("http://localhost:3128")
	if err == nil {
		t.Error("Failed to detect proper transport")
	}
}
