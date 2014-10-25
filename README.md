## Mocks

Mocks is a small package that provides mocks to help with testing network applications.

#### Mocking a network connection

To mock the `net.Conn` interface, import `github.com/gbbr/mocks` into your package
and use a configured mock structure, such as:

```go
var mockConn net.Conn

mockConn = &mocks.Conn{
	// Local address setup
	LAddr: "127.0.0.1:888",
	LNet:  "tcp",

	// Remote address
	RAddr: "10.18.20.21:123",
	RNet: "udp",
}

fmt.Println(mockConn.LocalAddr().String()) // prints "127.0.0.1:888"
fmt.Println(mockConn.RemoteAddr().String()) // prints "10.18.20.21:123"
```

The view data that was sent to the mock connection, configure the `In` io.Writer
interface of mocks.Conn, like:

```go
var buf bytes.Buffer
mockConn.In = &buf

fmt.Fprintf(mockConn, "Message")
fmt.Println(buf.String()) // prints "Message"
```

To set a data source for the network connection the `Out` io.Reader may be used as follows:

```go
var msg string

mockConn.Out = bytes.NewBuffer([]byte("Test\n"))
fmt.Scanln(mockConn, &msg)

fmt.Println(msg) // outputs "Test"
```


#### Obtaining a full communication channel

Pipe returns a full duplex network connection that receives data on either end and outputs
it on the other one. This functionality is similar to [net.Pipe](http://golang.org/pkg/net/#Pipe), but
additionally allows the mocking of addresses of each end using the connection from this package.

```go
c1, c2 := Pipe(
	&Conn{RAddr: "1.1.1.1:123"},
	&Conn{LAddr: "127.0.0.1:12", RAddr: "2.2.2.2:456"},
)

// Go routine writes to connection 1
go c1.Write([]byte("Hello"))

// Read 5 bytes
b := make([]byte, 5)

// Connection 2 receives message
n, err := c2.Read(b)
if err != nil {
	t.Errorf("Could not read c2: %s", err)
}

fmt.Println(string(b)) // outputs "Hello"
```

Refer to the [tests](https://github.com/gbbr/mocks/blob/master/conn_test.go#L75) for a complete example.

### Considerations

If you do not wish to to create the above examples (ie. you do not need to fake the remote/local address), you may also consider using the [pipe](http://golang.org/pkg/net/#Pipe) provided in the `net` package, which returns two ends of a network stream. _Careful though_, when using net.Pipe() and requesting LocalAddr() or RemoteAddr() nil pointer panic will happen.
