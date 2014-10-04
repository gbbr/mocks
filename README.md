## Go Mocks

A collection of mocks for testing Go applications.

### Usage

Mocking `net.Conn` has never been easier. Fake your Local & Remote IP address and protocol, channel incoming and outgoing stream of data as desired.

To fake your address & protocol:

```
var conn net.Conn = &mock.Conn{
  LocalAddress: "1.2.3.4:567",
  LocalNetwork: "tcp",
  
  RemoteAddress: "some.addr:666",
  RemoteNetwork: "udp",
}

fmt.Println(conn.LocalAddr())            // 1.2.3.4:567
fmt.Println(conn.RemoteAddr().Network()) // udp
```

Using a `net/textproto` wrapper is as easy as:

```
var conn net.Conn = new(mock.Conn)
var text = textproto.NewConn(conn)

text.PrintfLine("Hello world!")
fmt.Println(conn.Incoming.String()) // "Hello world!\r\n"
```

Check source code for more documentation. The mock interface is implemented as follows:

```go
type Conn struct {
	// Local network & address for the connection
	LocalNetwork, LocalAddress string

	// Remote network & address for the connection
	RemoteNetwork, RemoteAddress string

	// Incoing messages will be written to this buffer
	Incoming bytes.Buffer

	// Outgoing messages will be read from this buffer
	Outgoing *bytes.Buffer
}
```

### Considerations

If you do not wish to to create the above examples (ie. you do not need to fake the remote/local address), you may also consider using the [pipe](http://golang.org/pkg/net/#Pipe) provided in the `net` package, which returns two ends of a network stream.
