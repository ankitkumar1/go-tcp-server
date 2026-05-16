# TCP Server (Go Lang)
===

## Single Threaded : single_thread_server.go

## How to Run

```
$ go run single_thread_server.go
```

## How to connect from client side: 

```
$ nc localhost 7379
```

### Summary:
```
Listen()
This function sets up a network listener on a specific IP address and port, telling the operating system that your application is ready to receive incoming traffic. It acts like a receptionist opening a front desk and waiting for visitors to arrive.

Accept()
Accept() is a blocking call. When your code reaches this line, it completely pauses and waits until a client attempts to connect; once a client arrives, it unblocks, establishes a dedicated network connection (net.Conn), and hands it to you so you can talk to that specific client.

Read()
Read() is also a blocking call that waits for the connected client to send data over the network socket. It pauses execution until there are bytes available to read into your buffer, or until the client disconnects (which returns an io.EOF error).

Write()
Yes, Write() directly sends data back to the client over the established network socket. In your code, c.Write([]byte(cmd)) takes the string data you received, converts it to bytes, and pushes it right back down the open pipeline to the client's screen.

Is it a WebSocket connection?
No, this is a raw TCP connection, not a WebSocket.

WebSockets are a specialized protocol built on top of TCP that start with an HTTP request and are designed for web browsers. What you have built here is a lower-level, raw TCP Socket connection.

It stays alive because of your nested loop structure (for { ... }). In raw TCP, once a connection is accepted, the operating system keeps the underlying pipe open. Your code takes advantage of this by staying inside that inner loop, repeatedly Reading and Writeing to the same client until they explicitly disconnect or a network error occurs.

```

===
# Multi Threaded : multi_thread_tcp.go

## How to Run

```
$ go run multi_thread_server.go
```

## How to connect from client side: Fire this command from multiple terminals and send some messages.

```
$ nc localhost 1729
```