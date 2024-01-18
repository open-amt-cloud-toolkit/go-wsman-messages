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
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// WsTransport is an implementation of http.Transport which uses websocket relay
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
	buf_mutex sync.Mutex
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
		buf_mutex: sync.Mutex{},
	}
	return t
}

func (t *WsTransport) timedReadMessage(ms int) (b []byte) {
	timer := time.NewTimer(time.Duration(ms) * time.Millisecond)
	<-timer.C
	t.buf_mutex.Lock()
	b = append(b, t.messages...)
	t.messages = []byte{}
	t.buf_mutex.Unlock()
	return b
}

func (t *WsTransport) buildUrl() (url string) {
	// craft websocket url
	url = t.wsurl + "?p=" + strconv.Itoa(t.protocol) + "&host=" + t.host
	url += "&user=" + t.username + "&pass=" + t.password
	url += "&port=" + strconv.Itoa(t.port)
	if t.tls {
		url += "&tls=1"
	} else {
		url += "&tls=0"
	}
	if t.tls1only {
		url += "&tls1only=1"
	} else {
		url += "&tls1only=0"
	}
	return url
}

func (t *WsTransport) connectWebsocket() (conn *websocket.Conn, err error) {
	url := t.buildUrl()
	// Attempt to establish websocket connection
	var hdr = http.Header{}
	if t.token != "" {
		hdr.Set("Sec-Websocket-Protocol", t.token)
	}
	wsdialer := websocket.Dialer{}
	wsdialer.TLSClientConfig = t.tlsconfig
	conn, _, err = wsdialer.Dial(url, hdr)
	if err != nil {
		return nil, err
	} else {
		t.conn = conn
		go func() {
			for {
				//Trying to read.
				_, p, err := t.conn.ReadMessage()
				if err != nil {
					return
				}
				t.buf_mutex.Lock()
				t.messages = append(t.messages, p...)
				t.buf_mutex.Unlock()
			}
		}()
	}
	return conn, err
}

func (t *WsTransport) disconnectWebsocket() {
	if t.conn != nil {
		_ = t.conn.Close()
		t.conn = nil
	}
}

// RoundTrip makes a low level text exchange over websocket. This is supposed to be used by high level round tripper
func (t *WsTransport) RoundTrip(r *http.Request) (resp *http.Response, err error) {
	// Sanity check
	if t.wsurl == "" || t.protocol == 0 || t.host == "" || t.username == "" || t.password == "" || t.port == 0 {
		return nil, errors.New("Invalid Transport data")
	}
	// Check if we had already established websocket for this transport object, if not create
	if t.conn == nil || t.conn.UnderlyingConn() == nil {
		_, err := t.connectWebsocket()
		if err != nil {
			return nil, err
		}
	}
	// t.conn should be established
	// be careful when working with request Body.. make a copy
	buf, _ := io.ReadAll(r.Body)
	bd1 := io.NopCloser(bytes.NewBuffer(buf))
	bd2 := io.NopCloser(bytes.NewBuffer(buf))
	l, _ := io.Copy(io.Discard, bd1)
	r.Body = bd2
	r.Header.Add("Content-Length", strconv.FormatInt(l, 10))

	bytes_to_send, _ := httputil.DumpRequest(r, true)
	// write and ignore error status, proper error handling is at read go routine
	_ = t.conn.WriteMessage(websocket.TextMessage, bytes_to_send)

	b := t.timedReadMessage(100)
	bf := []byte{}
	bf = append(bf, b[:]...)
	// this is to check if we finished reading the whole body, otherwise read again
	// this will try to read data 10 times with 300 ms delay (3 seconds) if previous read yields empty data
	max_count := 10
	count := 0
	for {
		b = t.timedReadMessage(300)
		// how many zero read
		if len(b) == 0 {
			count = count + 1
			if count >= max_count {
				break
			}
		}
		bf = append(bf, b[:]...)
		if strings.Index(string(bf), "</a:Envelope>") > 0 {
			break
		}
		if strings.Index(string(bf), "</html>") > 0 {
			t.disconnectWebsocket()
			break
		}
	}
	read_buffer := bytes.NewReader(bf)
	resp, err = http.ReadResponse(bufio.NewReader(read_buffer), r)
	return resp, err
}
