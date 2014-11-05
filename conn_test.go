package mocks

import (
	"bytes"
	"fmt"
	"net"
	"net/textproto"
	"testing"
)

func Test_Interface_Is_Implemented(t *testing.T) {
	var (
		_ net.Conn = (*Conn)(nil)
		_ net.Addr = (*Addr)(nil)
	)
}

func Test_Conn_Incoming_Buffer(t *testing.T) {
	var buf bytes.Buffer

	mockConn := &Conn{In: &buf}

	fmt.Fprintln(mockConn, "Test Message")
	contents := buf.String()

	if contents != "Test Message\n" {
		t.Errorf("Expected 'Test Message' but got: '%s'", buf.String())
	}
}

func Test_Conn_Outgoing_Buffer(t *testing.T) {
	var contents = new(string)
	mockConn := &Conn{
		Out: bytes.NewBuffer([]byte("Test\n")),
	}

	fmt.Fscanln(mockConn, contents)

	if *contents != "Test" {
		t.Errorf("Expected 'Test' but got: '%s'", *contents)
	}
}

func Test_Conn_Textproto(t *testing.T) {
	var buf bytes.Buffer
	conn := &Conn{In: &buf}
	text := textproto.NewConn(conn)

	err := text.PrintfLine("Hello world!")
	if err != nil {
		t.Error("Could not write to connection.")
	}

	received := buf.String()
	if received != "Hello world!\r\n" {
		t.Errorf("Expected 'Hello world!', got: '%s'", received)
	}
}

func Test_Conn_Addresses(t *testing.T) {
	var conn net.Conn = &Conn{
		LNet:  "net1",
		RNet:  "net2",
		LAddr: "addr1",
		RAddr: "addr2",
	}

	if conn.LocalAddr().String() != "addr1" || conn.RemoteAddr().String() != "addr2" ||
		conn.LocalAddr().Network() != "net1" || conn.RemoteAddr().Network() != "net2" {
		t.Errorf("Did not mock addresses correctly")
	}
}

func Test_Pipe(t *testing.T) {
	c1, c2 := Pipe(
		&Conn{RAddr: "1.1.1.1:123"},
		&Conn{RAddr: "2.2.2.2:456"},
	)

	go c1.Write([]byte("Hello"))

	b := make([]byte, 5)
	n, err := c2.Read(b)
	if err != nil {
		t.Errorf("Could not read c2: %s", err)
	}

	if string(b) != "Hello" || n != 5 {
		t.Errorf("Pipe to c2 did not work, got %d bytes of '%s'", n, b)
	}

	if c1.RemoteAddr().String() != "1.1.1.1:123" || c2.RemoteAddr().String() != "2.2.2.2:456" {
		t.Error("Did not mock addresses correctly.")
	}

	go c2.Write([]byte("Jumbo"))

	n, err = c1.Read(b)
	if err != nil {
		t.Errorf("Could not read c1: %s", err)
	}

	if string(b) != "Jumbo" || n != 5 {
		t.Errorf("Pipe to c2 did not work, got %d bytes of '%s'", n, b)
	}
}
