package client

import (
	"crypto/tls"
	"fmt"
	"net"
	"sync"
)

func NewWsmanTCP(cp Parameters) *Target {
	port := RedirectionNonTLSPort
	if cp.UseTLS {
		port = RedirectionTLSPort
	}

	return &Target{
		endpoint:           cp.Target + ":" + port,
		username:           cp.Username,
		password:           cp.Password,
		useDigest:          cp.UseDigest,
		logAMTMessages:     cp.LogAMTMessages,
		challenge:          &AuthChallenge{},
		UseTLS:             cp.UseTLS,
		InsecureSkipVerify: cp.SelfSignedAllowed,
		bufferPool: sync.Pool{
			New: func() interface{} {
				return make([]byte, 4096) // Adjust size according to your needs.
			},
		},
	}
}

// Connect establishes a TCP connection to the endpoint specified in the Target struct.
func (t *Target) Connect() error {
	var err error
	if t.UseTLS {
		t.conn, err = tls.Dial("tcp", t.endpoint, &tls.Config{
			InsecureSkipVerify: t.InsecureSkipVerify,
		})
	} else {
		t.conn, err = net.Dial("tcp", t.endpoint)
	}

	if err != nil {
		return fmt.Errorf("failed to connect to %s: %w", t.endpoint, err)
	}

	return nil
}

// Send sends data to the connected TCP endpoint in the Target struct.
func (t *Target) Send(data []byte) error {
	if t.conn == nil {
		return fmt.Errorf("no active connection")
	}

	_, err := t.conn.Write(data)
	if err != nil {
		return fmt.Errorf("failed to send data: %w", err)
	}

	return nil
}

// Receive reads data from the connected TCP endpoint in the Target struct.
func (t *Target) Receive() ([]byte, error) {
	if t.conn == nil {
		return nil, fmt.Errorf("no active connection")
	}

	tmp := t.bufferPool.Get().([]byte)
	defer t.bufferPool.Put(tmp)

	n, err := t.conn.Read(tmp)
	if err != nil {
		return nil, err
	}

	return append([]byte(nil), tmp[:n]...), nil
}

// CloseConnection cleanly closes the TCP connection.
func (t *Target) CloseConnection() error {
	if t.conn == nil {
		return fmt.Errorf("no active connection to close")
	}

	err := t.conn.Close()
	if err != nil {
		return fmt.Errorf("failed to close connection: %w", err)
	}

	t.conn = nil

	return nil
}
