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

func TestAuthChallenge_parseChallenge(t *testing.T) {
	c := &authChallenge{}

	testCases := []struct {
		description string
		input       string
		expectQOP   string
	}{
		{
			description: "expect qop of auth for multiple options if auth is one of them",
			input:       `Digest realm="Digest:CA8E24F8F3E8EFBD8322A0BC846BDE7D", nonce="EzyJAVIHAAAAAAAA2hsF/jk27ByMWRKu",stale="false",qop="auth,auth-int, auth"`,
			expectQOP:   "auth",
		},
		{
			description: "expect qop of auth for multiple options in any order",
			input:       `Digest realm="Digest:CA8E24F8F3E8EFBD8322A0BC846BDE7D", nonce="EzyJAVIHAAAAAAAA2hsF/jk27ByMWRKu",stale="false",qop="auth-int,auth"`,
			expectQOP:   "auth",
		},
		{
			description: "expect qop of auth for single option",
			input:       `Digest realm="Digest:81D0A72A91B91FC7B08CED6F26DC0D49", nonce="kGuhOSX9AAAAAAAAjlLGwSwRfxiBgr2z",stale="false",qop="auth"`,
			expectQOP:   "auth",
		},
		{
			description: "expect qop of auth and ignore unkown options",
			input:       `Digest realm="Digest:81D0A72A91B91FC7B08CED6F26DC0D49", nonce="kGuhOSX9AAAAAAAAjlLGwSwRfxiBgr2z",stale="false",qop="auth,not-valid"`,
			expectQOP:   "auth",
		},
		{
			description: "expect qop equal to the original if auth not an option",
			input:       `Digest realm="Digest:CA8E24F8F3E8EFBD8322A0BC846BDE7D", nonce="EzyJAVIHAAAAAAAA2hsF/jk27ByMWRKu",stale="false",qop="auth-conf,auth-int"`,
			expectQOP:   "auth-conf,auth-int",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.description, func(t *testing.T) {
			err := c.parseChallenge(tc.input)
			assert.Nil(t, err)
			assert.Equal(t, tc.expectQOP, c.Qop)
		})

	}

}
