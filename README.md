## Go Mocks

A collection of mocks for testing Go applications. This was built as a necessity for [Gomez](https://github.com/gbbr/gomez), my Mail Exchange Server, to test Reverse Lookup success and failure.

### Usage

#### Mock a full duplex connection

If you need to use `net.Pipe` from the [net package](http://golang.org/pkg/net/#Pipe), but need a remote and local address, you can use the Pipe method provided in this package by passing it two mock connections, and in turn it will return a full-duplex pipe.

```go
	c1, c2 := Pipe(
		&Conn{RemoteAddress: "1.1.1.1:123"},
		&Conn{RemoteAddress: "2.2.2.2:456"},
	)
```

Refer to the [tests](https://github.com/gbbr/mocks/blob/master/conn_test.go#L75) for a complete example.

#### Mocking net.Conn

Mocking `net.Conn` has never been easier. Fake your Local & Remote IP address and protocol, channel incoming and outgoing stream of data as desired.

To fake your address & protocol:

```
var conn net.Conn 
var buf bytes.Buffer

conn = &mock.Conn{
  LocalAddress: "1.2.3.4:567",
  LocalNetwork: "tcp",
  
  RemoteAddress: "some.addr:666",
  RemoteNetwork: "udp",

  Incoming: &buf,
}

fmt.Println(conn.LocalAddr())            // 1.2.3.4:567
fmt.Println(conn.RemoteAddr().Network()) // udp

fmt.Fprintln(conn, "Message")
fmt.Println(buf.String()) // Outputs: Message\n
```

Using a `net/textproto` wrapper is as easy as:

```
var buf bytes.Buffer

conn := &mock.Conn{Incoming: &buf}
text := textproto.NewConn(conn)

text.PrintfLine("Hello world!")
fmt.Println(buf.String()) // "Hello world!\r\n"
```

Check source code for more documentation. The mock interface is implemented as follows:

```go
type Conn struct {
	// Local network & address for the connection
	LocalNetwork, LocalAddress string

	// Remote network & address for the connection
	RemoteNetwork, RemoteAddress string

	// Incoming messages will be written to this buffer
	Incoming io.Writer

	// Outgoing messages will be read from this buffer
	Outgoing io.Reader
}
```

### Considerations

If you do not wish to to create the above examples (ie. you do not need to fake the remote/local address), you may also consider using the [pipe](http://golang.org/pkg/net/#Pipe) provided in the `net` package, which returns two ends of a network stream. _Careful though_, when using net.Pipe() and requesting LocalAddr() or RemoteAddr() nil pointer panic will happen.
