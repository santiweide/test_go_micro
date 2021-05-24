package test

import (
	"context"
	"fmt"
	_ "github.com/mbobakov/grpc-consul-resolver" // It's important
	"github.com/micro/go-micro"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-micro/registry/consul"
	"test_go_micro"
	"test_go_micro/go_micro/model"
	"testing"
)

func BenchmarkTestString(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testString()
	}
}

func BenchmarkTestStruct(b *testing.B) {
	for n := 0; n < b.N; n++ {
		testStruct()
	}
}

func testString() {

	reg := consul.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = []string{
				"139.198.174.188:8500",
			}
		})

	// create a new service
	service := micro.NewService(micro.Registry(reg))

	// parse command line flags
	service.Init()
	greeter := model.NewGreeterService("Greeter", service.Client())
	req := &model.StringRequest{
		Message: test_go_micro.RandStringRunes(test_go_micro.Str10k),
	}

	// Use the generated client stub
	_, err := greeter.TestString(context.Background(), req)

	if err == nil {
		fmt.Printf("error:%v\n", err)
	}
}

func testStruct() {

	reg := consul.NewRegistry(
		func(options *registry.Options) {
			options.Addrs = []string{
				"139.198.174.188:8500",
			}
		})

	// create a new service
	service := micro.NewService(micro.Registry(reg))

	// parse command line flags
	service.Init()
	greeter := model.NewGreeterService("Greeter", service.Client())

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
	_, err := greeter.TestStruct(context.Background(), req)
	if err == nil {
		fmt.Printf("error:%v\n", err)
	}
}
