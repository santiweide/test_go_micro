package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	hello "github.com/micro/examples/greeter/srv/proto/hello"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
)

func main() {
	r := gin.Default()
	r.GET("/myRpc", func(c *gin.Context) {
		callServer()
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func callServer() {

	reg := consul.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = []string{
				"127.0.0.1:8500",
			}
		})

	// create a new service
	service := micro.NewService(micro.Registry(reg))

	// parse command line flags
	service.Init()

	// Use the generated client stub
	cl := hello.NewSayService("go.micro.srv.greeter", service.Client())

	_, err := cl.Hello(context.Background(), &hello.Request{
		Name: "John",
	})
	if err == nil {
		fmt.Printf("error:%v\n", err)

	}

}
