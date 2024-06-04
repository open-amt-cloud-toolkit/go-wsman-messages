/*********************************************************************
 * Copyright (c) Intel Corporation 2023
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/

package client

import (
	"bytes"
	"crypto/tls"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/url"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
)

const (
	ContentType           = "application/soap+xml; charset=utf-8"
	NSWSMAN               = "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"
	NSWSMID               = "http://schemas.dmtf.org/wbem/wsman/identity/1/wsmanidentity.xsd"
	TLSPort               = "16993"
	NonTLSPort            = "16992"
	RedirectionTLSPort    = "16995"
	RedirectionNonTLSPort = "16994"
)

type Message struct {
	XMLInput  string
	XMLOutput string
}

// WSMan is an interface for the wsman.Client.
type WSMan interface {
	// HTTP Methods
	Post(msg string) (response []byte, err error)
	// TCP Methods
	Connect() error
	Send(data []byte) error
	Receive() ([]byte, error)
	CloseConnection() error
}

// Target is a thin wrapper around http.Target.
type Target struct {
	http.Client
	endpoint       string
	username       string
	password       string
	useDigest      bool
	logAMTMessages bool
	challenge      *AuthChallenge
	conn           net.Conn
	bufferPool     sync.Pool
}

const timeout = 10 * time.Second

func NewWsman(cp Parameters) *Target {
	path := "/wsman"
	port := NonTLSPort

	if cp.UseTLS {
		port = TLSPort
	}

	protocol := "http"
	if port == TLSPort {
		protocol = "https"
	}

	res := &Target{
		endpoint:       protocol + "://" + cp.Target + ":" + port + path,
		username:       cp.Username,
		password:       cp.Password,
		useDigest:      cp.UseDigest,
		logAMTMessages: cp.LogAMTMessages,
	}

	res.Timeout = timeout
	res.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: cp.SelfSignedAllowed},
	}

	if res.useDigest {
		res.challenge = &AuthChallenge{Username: res.username, Password: res.password}
	}

	return res
}

// Post overrides http.Client's Post method.
func (t *Target) Post(msg string) (response []byte, err error) {
	msgBody := []byte(msg)

	var auth string

	bodyReader := bytes.NewReader(msgBody)

	req, err := http.NewRequest("POST", t.endpoint, bodyReader)
	if err != nil {
		return nil, err
	}

	if t.username != "" && t.password != "" {
		if t.useDigest {
			auth, err = t.challenge.authorize("POST", "/wsman")
			if err != nil {
				return nil, fmt.Errorf("failed digest auth %w", err)
			}

			if t.challenge.Realm != "" {
				req.Header.Set("Authorization", auth)
			}
		} else {
			req.SetBasicAuth(t.username, t.password)
		}
	}

	req.Header.Add("content-type", ContentType)

	if t.logAMTMessages {
		logrus.Trace(msg)
	}

	res, err := t.Do(req)
	if err != nil {
		return nil, err
	}

	if t.useDigest && res.StatusCode == 401 {
		if err := t.challenge.parseChallenge(res.Header.Get("WWW-Authenticate")); err != nil {
			return nil, err
		}

		auth, err = t.challenge.authorize("POST", "/wsman")
		if err != nil {
			return nil, fmt.Errorf("failed digest auth %w", err)
		}

		bodyReader = bytes.NewReader(msgBody)

		req, err = http.NewRequest("POST", t.endpoint, bodyReader)
		if err != nil {
			return nil, err
		}

		req.Header.Set("Authorization", auth)
		req.Header.Add("content-type", ContentType)

		res, err = t.Do(req)
		if err != nil && err.Error() != io.EOF.Error() {
			return nil, err
		}
	}

	defer res.Body.Close()

	if res.StatusCode >= 400 {
		b, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}

		if t.logAMTMessages {
			logrus.Trace(string(b))
		}

		errPostResponse := errors.New("wsman.Client post received")

		return nil, fmt.Errorf("%w: %v\n%v", errPostResponse, res.Status, string(b))
	}

	response, err = io.ReadAll(res.Body)

	if t.logAMTMessages {
		logrus.Trace(string(response))
	}

	if err != nil && err.Error() != io.EOF.Error() {
		return nil, err
	}

	return response, nil
}

// ProxyURL sets proxy address for the underlying Transport if supported.
func (t *Target) ProxyURL(proxyStr string) (err error) {
	// check if c.Transport is *http.Transport, otherwise currently it is not supported
	_, ok := t.Transport.(*http.Transport)
	if !ok {
		return errors.New("transport does not support proxy")
	}

	// check if proxy parsing failed or check if scheme is not nil
	proxyURL, err := url.Parse(proxyStr)
	if err != nil || (proxyURL != nil && proxyURL.Scheme == "") {
		return errors.New("unknown URL Scheme")
	}

	t.Transport.(*http.Transport).Proxy = http.ProxyURL(proxyURL)

	return nil
}
