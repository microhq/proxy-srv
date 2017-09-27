# Proxy SRV

The proxy-srv is a go-micro service which acts as a proxy for http applications. 

## Why is this useful?

How do you integrate legacy http applications into the micro ecosystem? The proxy-srv allows you to make RPC requests 
via your go-micro service to a legacy http application. The proxy-srv acts as a gateway to your http apps while 
allowing all the features of go-micro to be leveraged.

## How does it work?

Run a proxy-srv per http app in a 1:1 configuration. The proxy-srv should be registered under a pseudo name for the app. 
Just like go-micro services this will create an instance of the proxy per app you run. It acts as an RPC to HTTP proxy.

## Usage

Run discovery

```
consul agent -dev
```

Run the proxy-srv, give it a name and specify the app address you want to proxy for

Here we'll say it should be called com.example.srv.hello and we're proxying whatever runs on localhost:8089

```
proxy-srv --server_name=com.example.srv.hello --proxy_address=localhost:8089
```

Now call it like you would any go-micro service, making sure to use the Request/Response in the proto dir.

```
// create the proxy client
c := proto.NewProxyClient("com.example.srv.hello", cli.Client())

// make a request to "/hello"
rsp, err := c.Call(context.TODO(), &proto.Request{
	Method: "GET",
	Path:   "/hello",
})
```
