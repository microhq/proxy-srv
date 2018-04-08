package main

import (
	"fmt"

	"github.com/micro/go-micro"
	proto "github.com/micro/proxy-srv/proto"
	"golang.org/x/net/context"
)

func main() {
	cli := micro.NewService()
	cli.Init()

	// create the proxy client
	c := proto.ProxyServiceClient("com.example.srv.hello", cli.Client())

	// make the request
	rsp, err := c.Call(context.TODO(), &proto.Request{
		Method: "GET",
		Path:   "/hello",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	// print response
	fmt.Printf("got response %+v\n", rsp)
}
