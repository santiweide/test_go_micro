package main

import (
	"context"
	_ "github.com/mbobakov/grpc-consul-resolver"
	"google.golang.org/grpc"
	"log"
	"test_go_micro"
	"test_go_micro/grpc/model"
	"time"
)

const (
	target = "consul://127.0.0.1:8500/test_grpc"
)

func main() {
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

	// Contact the server and print out its response.
	//for {
	ctx, _ := context.WithTimeout(context.Background(), time.Second)
	_, err = c.TestString(ctx, &model.StringRequest{Message: test_go_micro.RandStringRunes(100)})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	//log.Printf("Greeting: %s", r.Message)
	//time.Sleep(time.Second * 2)
	//}
}
