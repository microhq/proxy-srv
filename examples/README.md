# Example

This example demonstrates the simple flow of a go-micro client calling the proxy via the name com.example.srv.hello and getting a plain 
text response from a http based application.

[go-micro client] => proxy-srv => http app

## Usage

Run discovery

```
consul agent -dev
```

Run the proxy with the name com.example.srv.hello and proxy address set to the apps address localhost:8089

```
proxy-srv --server_name=com.example.srv.hello --proxy_address=localhost:8089
```

Run the app

```
go run app/main.go
```

Run the client

```
go run cli/main.go
```

Expected response

```
got response statusCode:200 header:<key:"Content-Length" value:<key:"Content-Length" values:"11" > > header:<key:"Content-Type" value:<key:"Content-Type" values:"text/plain; charset=utf-8" > > header:<key:"Date" value:<key:"Date" values:"Wed, 27 Sep 2017 16:40:14 GMT" > > body:"hello world" 
```
