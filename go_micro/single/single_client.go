package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"log"
	"test_go_micro"
	"test_go_micro/go_micro/model"
	"unsafe"
)

func main() {
	testString1K()
}


func testString1K() {

	reg := consul.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = []string{
				"192.168.0.3:8500",
			}
		})

	// create a new service
	service := micro.NewService(micro.Registry(reg))

	// parse command line flags
	service.Init()
	greeter := model.NewGreeterService("Greeter", service.Client())
	req := &model.StringRequest{
		Message: test_go_micro.RandStringRunes(test_go_micro.Str1k),
	}

	log.Printf("Request Size: %v\n", unsafe.Sizeof(req))
	// Use the generated client stub
	resp, err := greeter.TestString(context.Background(), req)

	if err == nil {
		fmt.Printf("error:%v\n", err)
	} else {
		fmt.Printf("resp: %v\n", resp)
	}
}

func testStruct() {

	reg := consul.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = []string{
				"192.168.0.3:8500",
			}
		})

	// create a new service
	service := micro.NewService(micro.Registry(reg))

	// parse command line flags
	service.Init()
	greeter := model.NewGreeterService("Greeter", service.Client())

	_map := make(map[int32]string)
	_list := make([]string,100)
	for i := 0;i < 100;i ++ {
		_map[int32(i)]=test_go_micro.RandStringRunes(100)
		_list[i] = test_go_micro.RandStringRunes(100)
	}
	req := &model.StructRequest{
		Id:         123,
		KvMap:      _map,
		StringList: _list,
	}
	log.Printf("Request Size: %v\n", unsafe.Sizeof(req))
	// Use the generated client stub
	resp, err := greeter.TestStruct(context.Background(), req)
	if err == nil {
		fmt.Printf("error:%v\n", err)
	} else {
		fmt.Println("resp: %v\n", resp)
	}
}