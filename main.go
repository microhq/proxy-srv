package main

import (
	"strings"

	"github.com/micro/cli"
	"github.com/micro/go-log"
	"github.com/micro/go-micro"
	"github.com/microhq/proxy-srv/handler"
	proto "github.com/microhq/proxy-srv/proto"
)

func main() {
	proxyHandler := new(handler.Proxy)

	service := micro.NewService(
		micro.Name("go.micro.srv.proxy"),
		micro.Flags(
			cli.StringFlag{
				Name:  "proxy_address",
				Usage: "Specify the address to be proxied. Formats accepted host:port, http://host:port, https://host:port",
			},
		),
		micro.Action(func(c *cli.Context) {
			proxyAddr := c.String("proxy_address")

			if len(proxyAddr) == 0 {
				log.Fatal("Set the --proxy_address flag to the host to be proxied")
			}

			switch {
			case strings.HasPrefix(proxyAddr, "http://"):
				proxyHandler.Url = proxyAddr
			case strings.HasPrefix(proxyAddr, "https://"):
				proxyHandler.Url = proxyAddr
			default:
				proxyHandler.Url = "http://" + proxyAddr
			}
		}),
	)

	service.Init()

	// set our internal name for errors
	proxyHandler.Name = service.Server().Options().Name

	proto.RegisterProxyHandler(service.Server(), proxyHandler)

	if err := service.Run(); err != nil {
		log.Fatal(err)
	}
}
