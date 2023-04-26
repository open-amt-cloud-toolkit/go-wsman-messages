package wsman

import (
	"bytes"
	"crypto/tls"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

const ContentType = "application/soap+xml; charset=utf-8"
const NS_WSMAN = "http://schemas.dmtf.org/wbem/wsman/1/wsman.xsd"
const NS_WSMID = "http://schemas.dmtf.org/wbem/wsman/identity/1/wsmanidentity.xsd"

// Client is a thin wrapper around http.Client.
type Client struct {
	http.Client
	endpoint     string
	username     string
	password     string
	useDigest    bool
	OptimizeEnum bool
	challenge    *authChallenge
}

func NewClient(target, username, password string, useDigest bool) *Client {
	res := &Client{
		endpoint:  target,
		username:  username,
		password:  password,
		useDigest: useDigest,
	}
	res.Timeout = 10 * time.Second
	res.Transport = &http.Transport{
		TLSClientConfig: &tls.Config{},
	}
	if res.useDigest {
		res.challenge = &authChallenge{Username: res.username, Password: res.password}
		resp, err := res.PostForm(res.endpoint, nil)
		if err != nil {
			log.Fatalf("unable to perform digest auth with %s: %v", res.endpoint, err)
		}
		if resp.StatusCode != 401 {
			log.Fatalf("no digest auth at %s", res.endpoint)
		}
		if err := res.challenge.parseChallenge(resp.Header.Get("WWW-Authenticate")); err != nil {
			log.Fatalf("failed to parse auth header %v", err)
		}
	}
	return res
}

// Post overrides http.Client's Post method
func (c *Client) Post(msg string) (response []byte, err error) {

	msgBody := []byte(msg)
	bodyReader := bytes.NewReader(msgBody)
	req, err := http.NewRequest("POST", c.endpoint, bodyReader)
	if err != nil {
		return nil, err
	}
	fmt.Printf("req:%#v\nbody:\n%s\n", req, msg)
	if c.username != "" && c.password != "" {
		if c.useDigest {
			auth, err := c.challenge.authorize("POST", c.endpoint)
			if err != nil {
				return nil, fmt.Errorf("failed digest auth %v", err)
			}
			req.Header.Set("Authorization", auth)
		} else {
			req.SetBasicAuth(c.username, c.password)
		}
	}
	req.Header.Add("content-type", ContentType)

	res, err := c.Do(req)
	if err != nil {
		return nil, err
	}
	if c.useDigest && res.StatusCode == 401 {
		log.Print("Digest reauthorizing")
		if err := c.challenge.parseChallenge(res.Header.Get("WWW-Authenticate")); err != nil {
			return nil, err
		}
		auth, err := c.challenge.authorize("POST", "/wsman")
		if err != nil {
			return nil, fmt.Errorf("failed digest auth %v", err)
		}
		bodyReader = bytes.NewReader(msgBody)
		req, err = http.NewRequest("POST", c.endpoint, bodyReader)
		if err != nil {
			return nil, err
		}
		req.Header.Set("Authorization", auth)
		req.Header.Add("content-type", ContentType)
		res, err = c.Do(req)
		if err != nil {
			return nil, err
		}
	}

	defer res.Body.Close()

	if res.StatusCode >= 400 {
		b, _ := io.ReadAll(res.Body)
		return nil, fmt.Errorf("wsman.Client: post recieved %v\n'%v'", res.Status, string(b))
	}
	response, err = io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return response, nil
}
