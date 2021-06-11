package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"log"
	"test_go_micro/go_micro/model"
	"time"
)

type Greeter struct{}

func (g *Greeter) TestString(ctx context.Context, req *model.StringRequest, resp *model.StringResponse) error {
	log.Printf("reqMessage: %v", req.Message)
	resp.Message = req.Message
	return nil
}

func (g *Greeter) TestStruct(ctx context.Context, req *model.StructRequest, resp *model.StructResponse) error {
	resp.Id = req.Id
	resp.KvMap = req.KvMap
	resp.StringList = req.StringList
	return nil
}

type HelloService struct{}

func (h *HelloService) Hello(ctx context.Context, req *model.HelloRequest,resp *model.HelloResponse)  error{
	resp.Greeting = " 你好, " + req.Name
	return nil
}

func main() {
	reg := consul.NewRegistry(
		func(options *registry.Options){
			options.Addrs = []string{
				"192.168.0.3:8500",
			}
		})

	// 创建新的服务
	service := micro.NewService(
		micro.Registry(reg),
		micro.Name("com.dut.srv.greeter"),
		micro.RegisterTTL(time.Second*30000),
		micro.RegisterInterval(time.Second*100),
	)

	// 初始化，会解析命令行参数
	service.Init()

	// 注册处理器，调用 Greeter 服务接口处理请求
	model.RegisterGreeterHandler(service.Server(), new(Greeter))
	model.RegisterHelloServerHandler(service.Server(), new(HelloService))
	// 启动服务
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
