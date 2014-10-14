package mocks

import (
	"bytes"
	"io"
	"net"
	"time"
)

// Mocks a network connection. Implements the net.Conn interface
type Conn struct {
	// Local network & address for the connection
	LocalNetwork, LocalAddress string

	// Remote network & address for the connection
	RemoteNetwork, RemoteAddress string

	// Incoming messages will be written to this buffer
	Incoming bytes.Buffer

	// Outgoing messages will be read from this buffer
	Outgoing io.Reader
}

func (c *Conn) Read(b []byte) (n int, err error) {
	return c.Outgoing.Read(b)
}

func (c *Conn) Write(b []byte) (n int, err error) {
	return c.Incoming.Write(b)
}

func (c Conn) LocalAddr() net.Addr {
	return &Addr{c.LocalNetwork, c.LocalAddress}
}

func (c Conn) RemoteAddr() net.Addr {
	return &Addr{c.RemoteNetwork, c.RemoteAddress}
}

// Not implemented
func (c Conn) Close() error                       { return nil }
func (c Conn) SetDeadline(t time.Time) error      { return nil }
func (c Conn) SetReadDeadline(t time.Time) error  { return nil }
func (c Conn) SetWriteDeadline(t time.Time) error { return nil }

// Mocks a network address
type Addr struct {
	Net, Addr string
}

func (m Addr) Network() string {
	return m.Net
}

func (m Addr) String() string {
	return m.Addr
}
