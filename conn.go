package mocks

import (
	"io"
	"net"
	"syscall"
	"time"
)

// Mocks a network connection. Implements the net.Conn interface
type Conn struct {
	// Local network & address for the connection
	LocalNetwork, LocalAddress string

	// Remote network & address for the connection
	RemoteNetwork, RemoteAddress string

	// Incoming messages will be written to this buffer
	Incoming io.Writer

	// Outgoing messages will be read from this buffer
	Outgoing io.Reader

	closed bool
}

func (c *Conn) Read(b []byte) (n int, err error) {
	if c.closed {
		return 0, syscall.EINVAL
	}

	return c.Outgoing.Read(b)
}

func (c *Conn) Write(b []byte) (n int, err error) {
	if c.closed {
		return 0, syscall.EINVAL
	}

	return c.Incoming.Write(b)
}

func (c Conn) LocalAddr() net.Addr {
	return &Addr{c.LocalNetwork, c.LocalAddress}
}

func (c Conn) RemoteAddr() net.Addr {
	return &Addr{c.RemoteNetwork, c.RemoteAddress}
}

func (c *Conn) Close() error {
	c.closed = true
	return nil
}

// Not implemented
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

// Pipe turns two mock connections into a full-duplex connection similar to net.Pipe
// to allow pipe's with (fake) addresses.
func Pipe(c1, c2 *Conn) (*Conn, *Conn) {
	r1, w1 := io.Pipe()
	r2, w2 := io.Pipe()

	c1.Incoming = w1
	c2.Outgoing = r1

	c1.Outgoing = r2
	c2.Incoming = w2

	return c1, c2
}
