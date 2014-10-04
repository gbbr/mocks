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
		_ net.Conn = &Conn{}
		_ net.Addr = &Addr{}
	)
}

func Test_Conn_Incoming_Buffer(t *testing.T) {
	mockConn := new(Conn)
	fmt.Fprintln(mockConn, "Test Message")

	contents := mockConn.Incoming.String()
	if contents != "Test Message\n" {
		t.Errorf("Expected 'Test Message' but got: '%s'", mockConn.Incoming.String())
	}
}

func Test_Conn_Outgoing_Buffer(t *testing.T) {
	var contents = new(string)
	mockConn := &Conn{
		Outgoing: *(bytes.NewBuffer([]byte("Test\n"))),
	}

	fmt.Fscanln(mockConn, contents)

	if *contents != "Test" {
		t.Errorf("Expected 'Test' but got: '%s'", *contents)
	}
}

func Test_Conn_Textproto(t *testing.T) {
	conn := &Conn{}
	text := textproto.NewConn(conn)

	err := text.PrintfLine("Hello world!")
	if err != nil {
		t.Error("Could not write to connection.")
	}

	received, err := conn.Incoming.ReadString('!')
	if received != "Hello world!" {
		t.Errorf("Expected 'Hello world!', got: '%s'", received)
	}
}

func Test_Conn_Addresses(t *testing.T) {
	var conn net.Conn = &Conn{
		LocalNetwork:  "net1",
		RemoteNetwork: "net2",
		LocalAddress:  "addr1",
		RemoteAddress: "addr2",
	}

	if conn.LocalAddr().String() != "addr1" || conn.RemoteAddr().String() != "addr2" ||
		conn.LocalAddr().Network() != "net1" || conn.RemoteAddr().Network() != "net2" {
		t.Errorf("Did not mock addresses correctly")
	}
}
