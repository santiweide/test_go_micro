package test

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"test_go_micro"
	"test_go_micro/grpc/model"
	"testing"
)

const (
	target = "139.198.174.188:8500/test_grpc"
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
	conn, err := grpc.Dial(
		target,
		//		"consul://127.0.0.1:8500/test_grpc",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := model.NewGreeterClient(conn)

	req := &model.StringRequest{
		Message: test_go_micro.RandStringRunes(10000),
	}

	_, err = c.TestString(context.Background(), req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}

func testStruct() {
	conn, err := grpc.Dial(
		target,
		//		"consul://127.0.0.1:8500/test_grpc",
		grpc.WithInsecure(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy": "round_robin"}`),
	)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	c := model.NewGreeterClient(conn)

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

	_, err = c.TestStruct(context.Background(), req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}
