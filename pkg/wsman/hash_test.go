/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package wsman

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashWithMD5(t *testing.T) {
	testCases := []struct {
		data     string
		expected string
	}{
		{"test", "098f6bcd4621d373cade4e832627b4f6"},
		{"", "d41d8cd98f00b204e9800998ecf8427e"},
	}

	for _, tc := range testCases {
		actual := hashWithMD5(tc.data)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestHashWithHash(t *testing.T) {
	testCases := []struct {
		secret   string
		data     string
		expected string
	}{
		{"test", "data", "6681a96ff3ec7263ba55963bc91f5c72"},
		{"", "", "853ae90f0351324bd73ea615e6487517"},
	}

	for _, tc := range testCases {
		actual := hashWithHash(tc.secret, tc.data)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestAuthChallenge_HashCredentials(t *testing.T) {
	c := &authChallenge{
		Username: "test",
		Realm:    "realm",
		Password: "pass",
	}

	expected := "cccddec7bd6f77524ddac0e981fe5ba8"
	actual := c.hashCredentials()
	assert.Equal(t, expected, actual)
}

func TestAuthChallenge_HashURI(t *testing.T) {
	c := &authChallenge{}
	testCases := []struct {
		method   string
		uri      string
		expected string
	}{
		{"GET", "/path", "006e982dfebea99c2ce000b38b68e162"},
		{"POST", "/path", "431ef981d77e339ff4cde79dc5f4a44d"},
	}

	for _, tc := range testCases {
		actual := c.hashURI(tc.method, tc.uri)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestResponse(t *testing.T) {
	c := &authChallenge{}
	testCases := []struct {
		method   string
		uri      string
		cnonce   string
		expected string
	}{
		{"POST", "/path", "1b2311bfa21bfd81", "b9c5d3509b8a70d95a6204668265e6f9"},
	}

	for _, tc := range testCases {
		actual, _ := c.response(tc.method, tc.uri, tc.cnonce)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestResponse_Qop(t *testing.T) {
	c := &authChallenge{}
	testCases := []struct {
		method   string
		uri      string
		cnonce   string
		expected string
	}{
		{"POST", "/path", "1b2311bfa21bfd81", "c6b417c8da94a1d3649b4f6894c78379"},
	}

	for _, tc := range testCases {
		c.Qop = "auth,auth-int, auth"
		actual, _ := c.response(tc.method, tc.uri, tc.cnonce)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestAuthorize(t *testing.T) {
	c := &authChallenge{}
	testCases := []struct {
		method   string
		uri      string
		expected string
	}{
		{"POST", "/path", "Digest username=\"\", realm=\"\", nonce=\"\", uri=\"/path\", response=\"b9c5d3509b8a70d95a6204668265e6f9\""},
	}

	for _, tc := range testCases {
		actual, _ := c.authorize(tc.method, tc.uri)
		assert.Equal(t, tc.expected, actual)
	}
}

func TestAuthorize_Full(t *testing.T) {
	c := &authChallenge{}
	testCases := []struct {
		method   string
		uri      string
		expected string
	}{
		{"POST", "/path", "Digest username=\"admin\", realm=\"test\", nonce=\"00000001\", uri=\"/path\", response=\"99d726faaaa4e8874b3fd0f425053f5a\""},
	}

	for _, tc := range testCases {
		c.Username = "admin"
		c.Realm = "test"
		c.Nonce = "00000001"
		actual, _ := c.authorize(tc.method, tc.uri)
		assert.Equal(t, tc.expected, actual)
	}
}
