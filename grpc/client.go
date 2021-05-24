package main

import (
	"context"
	"github.com/gin-gonic/gin"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"google.golang.org/grpc"
	"log"
	"test_go_micro"
	"test_go_micro/grpc/model"
	"time"
	"unsafe"
)

const (
	target = "consul://192.168.0.4:8500/test_grpc"
)

func main(){
	r := gin.Default()
	r.GET("/string", func(c *gin.Context) {
		testString()
	})
	r.GET("/struct", func(c *gin.Context) {
		testStruct()
	})
	r.Run() // listen and serve on 0.0.0.0:8080
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
		Message: test_go_micro.RandStringRunes(test_go_micro.StrLen),
	}

	log.Printf("Request Size: %v\n", unsafe.Sizeof(req))
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	_, err = c.TestString(ctx, req)
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

	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	_, err = c.TestStruct(ctx, req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
}
