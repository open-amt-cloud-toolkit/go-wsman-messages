/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package wsman

import (
	"crypto/md5"
	"crypto/rand"
	"fmt"
	"io"
	"log"
	"strings"
)

type authChallenge struct {
	Username   string
	Password   string
	Realm      string
	CSRFToken  string
	Domain     string
	Nonce      string
	Opaque     string
	Stale      string
	Algorithm  string
	Qop        string
	CNonce     string
	NonceCount int
}

func hashWithMD5(data string) string {
	md5Hash := md5.New()
	_, err := io.WriteString(md5Hash, data)
	if err != nil {
		log.Println("failed to write string to md5 hash")
	}
	return fmt.Sprintf("%x", md5Hash.Sum(nil))
}

func hashWithHash(secret, data string) string {
	return hashWithMD5(fmt.Sprintf("%s:%s", secret, data))
}

func (c *authChallenge) hashCredentials() string {
	return hashWithMD5(fmt.Sprintf("%s:%s:%s", c.Username, c.Realm, c.Password))
}

func (c *authChallenge) hashURI(method, uri string) string {
	return hashWithMD5(fmt.Sprintf("%s:%s", method, uri))
}

func (c *authChallenge) response(method, uri, cnonce string) (string, error) {
	c.NonceCount++

	if strings.Contains(c.Qop, "auth") || c.Qop == "" {
		nonceData := c.Nonce
		if strings.Contains(c.Qop, "auth") {
			if cnonce != "" {
				c.CNonce = cnonce
			} else {
				b := make([]byte, 8)
				if _, err := io.ReadFull(rand.Reader, b); err != nil {
					return "", fmt.Errorf("failed to generate random bytes: %v", err)
				}
				c.CNonce = fmt.Sprintf("%x", b)[:16]
			}
			c.Qop = "auth"
			nonceData = fmt.Sprintf("%s:%08x:%s:%s", nonceData, c.NonceCount, c.CNonce, c.Qop)
		}

		hashedCredentials := c.hashCredentials()
		hashedURI := c.hashURI(method, uri)
		response := hashWithHash(hashedCredentials, fmt.Sprintf("%s:%s", nonceData, hashedURI))

		return response, nil
	}

	return "", fmt.Errorf("not implemented")
}

func (c *authChallenge) authorize(method, uri string) (string, error) {

	if !strings.Contains(c.Qop, "auth") && c.Qop != "" {
		return "", fmt.Errorf("qop not implemented")
	}
	response, err := c.response(method, uri, "")
	if err != nil {
		return "", err
	}

	var sb strings.Builder
	sb.WriteString(`Digest username="`)
	sb.WriteString(c.Username)
	sb.WriteString(`", realm="`)
	sb.WriteString(c.Realm)
	sb.WriteString(`", nonce="`)
	sb.WriteString(c.Nonce)
	sb.WriteString(`", uri="`)
	sb.WriteString(uri)
	sb.WriteString(`", response="`)
	sb.WriteString(response)
	sb.WriteString(`"`)

	if c.Algorithm != "" {
		sb.WriteString(`, algorithm="`)
		sb.WriteString(c.Algorithm)
		sb.WriteString(`"`)
	}
	if c.Opaque != "" {
		sb.WriteString(`, opaque="`)
		sb.WriteString(c.Opaque)
		sb.WriteString(`"`)
	}
	if c.Qop != "" {
		sb.WriteString(`, qop="`)
		sb.WriteString(c.Qop)
		sb.WriteString(`", nc="`)
		sb.WriteString(fmt.Sprintf("%08x", c.NonceCount))
		sb.WriteString(`", cnonce="`)
		sb.WriteString(c.CNonce)
		sb.WriteString(`"`)
	}

	return sb.String(), nil
}

func (c *authChallenge) parseChallenge(input string) error {
	const ws = " \n\r\t"
	const qs = "\""
	s := strings.Trim(input, ws)
	if !strings.HasPrefix(s, "Digest ") {
		return fmt.Errorf("challenge is bad, missing digest prefix: %s", input)
	}
	s = strings.Trim(s[7:], ws)
	sl := strings.Split(s, "\",")
	c.Algorithm = "MD5"
	var r []string
	for _, elem := range sl {
		r = strings.SplitN(elem, "=", 2)
		if len(r) != 2 {
			return fmt.Errorf("challenge is bad, malformed token: %s", elem)
		}
		key := strings.TrimSpace(r[0])
		value := strings.Trim(strings.TrimSpace(r[1]), qs)
		switch key {
		case "realm":
			c.Realm = value
		case "domain":
			c.Domain = value
		case "nonce":
			c.Nonce = value
		case "opaque":
			c.Opaque = value
		case "stale":
			c.Stale = value
		case "algorithm":
			c.Algorithm = value
		case "qop":
			c.Qop = value
		default:
			return fmt.Errorf("challenge is bad, unexpected token: %s", sl)
		}
	}
	return nil
}
