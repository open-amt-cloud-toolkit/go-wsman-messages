/*********************************************************************
 * Copyright (c) Intel Corporation 2024
 * SPDX-License-Identifier: Apache-2.0
 **********************************************************************/
package client

import (
	"bufio"
	"bytes"
	"crypto/tls"
	"errors"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// WsTransport is an implementation of http.Transport which uses websocket relay.
type WsTransport struct {
	wsurl     string
	protocol  int
	host      string
	username  string
	password  string
	port      int
	tls       bool
	tls1only  bool
	token     string
	conn      *websocket.Conn
	tlsconfig *tls.Config
	bufMutex  sync.Mutex
	messages  []byte
}

// NewTransport creates a new Websocket RoundTripper.
func NewWsTransport(wsurl string, protocol int, host, username, password string, port int, tls, tls1only bool, token string, tlsconfig *tls.Config) *WsTransport {
	t := &WsTransport{
		wsurl:     wsurl,
		protocol:  protocol,
		host:      host,
		username:  username,
		password:  password,
		port:      port,
		tls:       tls,
		tls1only:  tls1only,
		token:     token,
		tlsconfig: tlsconfig,
		bufMutex:  sync.Mutex{},
	}

	return t
}

func (t *WsTransport) timedReadMessage(ms int) (b []byte) {
	timer := time.NewTimer(time.Duration(ms) * time.Millisecond)
	<-timer.C
	t.bufMutex.Lock()
	b = append(b, t.messages...)
	t.messages = []byte{}
	t.bufMutex.Unlock()

	return b
}

func (t *WsTransport) buildURL() string {
	// Use net/url to construct the URL
	u, err := url.Parse(t.wsurl)
	if err != nil {
		return ""
	}

	// Prepare query parameters
	q := u.Query()
	q.Set("p", strconv.Itoa(t.protocol))
	q.Set("host", t.host)
	q.Set("user", t.username)
	q.Set("pass", t.password)
	q.Set("port", strconv.Itoa(t.port))
	q.Set("tls", strconv.FormatBool(t.tls))
	q.Set("tls1only", strconv.FormatBool(t.tls1only))
	// Set query string to URL
	u.RawQuery = q.Encode()

	return u.String()
}

func (t *WsTransport) connectWebsocket() (conn *websocket.Conn, err error) {
	url := t.buildURL()

	// Attempt to establish websocket connection
	hdr := http.Header{}
	if t.token != "" {
		hdr.Set("Sec-Websocket-Protocol", t.token)
	}

	wsdialer := websocket.Dialer{}
	wsdialer.TLSClientConfig = t.tlsconfig

	conn, _, err = wsdialer.Dial(url, hdr)
	if err != nil {
		return nil, err
	}

	t.conn = conn

	go func() {
		for {
			// Trying to read.
			var p []byte

			_, p, err = t.conn.ReadMessage()
			if err != nil {
				return
			}

			t.bufMutex.Lock()

			t.messages = append(t.messages, p...)

			t.bufMutex.Unlock()
		}
	}()

	return conn, err
}

func (t *WsTransport) disconnectWebsocket() {
	if t.conn != nil {
		_ = t.conn.Close()
		t.conn = nil
	}
}

// RoundTrip makes a low level text exchange over websocket. This is supposed to be used by high level round tripper.
func (t *WsTransport) RoundTrip(r *http.Request) (resp *http.Response, err error) {
	// Sanity check
	if t.wsurl == "" || t.protocol == 0 || t.host == "" || t.username == "" || t.password == "" || t.port == 0 {
		return nil, errors.New("invalid transport data")
	}

	// Check if we had already established websocket for this transport object, if not create
	if t.conn == nil || t.conn.UnderlyingConn() == nil {
		_, err = t.connectWebsocket()
		if err != nil {
			return nil, err
		}
	}

	// t.conn should be established
	// be careful when working with request Body.. make a copy
	buf, err := io.ReadAll(r.Body)
	if err != nil {
		return nil, err
	}

	bd1 := io.NopCloser(bytes.NewBuffer(buf))
	bd2 := io.NopCloser(bytes.NewBuffer(buf))

	l, err := io.Copy(io.Discard, bd1)
	if err != nil {
		return nil, err
	}

	r.Body = bd2
	r.Header.Add("Content-Length", strconv.FormatInt(l, 10))

	bytesToSend, err := httputil.DumpRequest(r, true)
	if err != nil {
		return nil, err
	}

	// write and ignore error status, proper error handling is at read go routine
	err = t.conn.WriteMessage(websocket.TextMessage, bytesToSend)
	if err != nil {
		return nil, err
	}

	b := t.timedReadMessage(100)
	bf := []byte{}
	bf = append(bf, b...)
	// this is to check if we finished reading the whole body, otherwise read again
	// this will try to read data 10 times with 300 ms delay (3 seconds) if previous read yields empty data
	maxCount := 10
	count := 0

	for {
		b = t.timedReadMessage(300)
		// how many zero read
		if len(b) == 0 {
			count++
			if count >= maxCount {
				break
			}
		}

		bf = append(bf, b...)
		if bytes.Index(bf, []byte("</a:Envelope>")) > 0 {
			break
		}

		if bytes.Index(bf, []byte("</html>")) > 0 {
			t.disconnectWebsocket()

			break
		}
	}

	readBuffer := bytes.NewReader(bf)
	resp, err = http.ReadResponse(bufio.NewReader(readBuffer), r)

	return resp, err
}
