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
	LNet, LAddr string
	// Remote network & address for the connection
	RNet, RAddr string

	// In messages will be written to this buffer
	In io.Writer
	// Out messages will be read from this buffer
	Out io.Reader

	closed bool
}

func (c *Conn) Read(b []byte) (n int, err error) {
	if c.closed {
		return 0, syscall.EINVAL
	}

	return c.Out.Read(b)
}

func (c *Conn) Write(b []byte) (n int, err error) {
	if c.closed {
		return 0, syscall.EINVAL
	}

	return c.In.Write(b)
}

func (c Conn) LocalAddr() net.Addr {
	return &Addr{c.LNet, c.LAddr}
}

func (c Conn) RemoteAddr() net.Addr {
	return &Addr{c.RNet, c.RAddr}
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

func (m Addr) Network() string { return m.Net }
func (m Addr) String() string  { return m.Addr }

// Pipe turns two mock connections into a full-duplex connection similar to net.Pipe
// to allow pipe's with (fake) addresses.
func Pipe(c1, c2 *Conn) (*Conn, *Conn) {
	r1, w1 := io.Pipe()
	r2, w2 := io.Pipe()

	c1.In = w1
	c2.Out = r1

	c1.Out = r2
	c2.In = w2

	return c1, c2
}
