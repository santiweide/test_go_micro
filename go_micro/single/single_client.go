package main

import (
	"context"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"log"
	"test_go_micro"
	"test_go_micro/go_micro/model"
)

func main() {
	testHello()
	testString1K()
	testStruct()
}

func testHello() {
	reg := consul.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = []string{
				"127.0.0.1:8500",
			}
		})
	service := micro.NewService(micro.Registry(reg), micro.Name("HelloServer.Client"))

	// 初始化
	service.Init()

	// 创建 Greeter 客户端
	greeter := model.NewHelloServerService("com.dut.srv.greeter", service.Client())

	// 远程调用 Greeter 服务的 Hello 方法
	rsp, err := greeter.Hello(context.TODO(), &model.HelloRequest{Name: "DUT"})
	if err != nil {
		log.Printf("error: %v\n", err)
	}
	if rsp == nil {
		log.Printf("resp is nil")
	} else {
		// Print response
		log.Printf("resp: %v\n", rsp.Greeting)
	}
}

func testString1K() {

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
	greeter := model.NewGreeterService("com.dut.srv.greeter", service.Client())
	req := &model.StringRequest{
		Message: test_go_micro.RandStringRunes(test_go_micro.Str1k),
	}

	// Use the generated client stub
	resp, err := greeter.TestString(context.Background(), req)

	if err != nil {
		log.Printf("error:%v\n", err)
	}
	if resp != nil {
		log.Printf("resp: %v\n", resp.Message)
	} else {
		log.Printf("resp is nil")
	}
}

func testStruct() {

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
	greeter := model.NewGreeterService("com.dut.srv.greeter", service.Client())

	_map := make(map[int32]string)
	_list := make([]string, 100)
	for i := 0; i < 100; i++ {
		_map[int32(i)] = test_go_micro.RandStringRunes(100)
		_list[i] = test_go_micro.RandStringRunes(100)
	}
	req := &model.StructRequest{
		Id:         123,
		KvMap:      _map,
		StringList: _list,
	}

	// Use the generated client stub
	resp, err := greeter.TestStruct(context.Background(), req)
	if err != nil {
		log.Printf("error:%v\n", err)
	}
	if resp != nil {
		log.Printf("resp: %v  \n", resp)
	} else {
		log.Printf("resp is nil")
	}
}
