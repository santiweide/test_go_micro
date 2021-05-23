package main

import (
	"context"
	"github.com/micro/go-micro/registry/consul"
	"google.golang.org/grpc"
	"log"
	"os"
	"test_go_micro/grpc/model"
	"time"
)

const (
	target      = "consul://127.0.0.1:8500/helloworld"
	defaultName = "world"
)

func main() {
	consul.Init()
	// Set up a connection to the server.
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	conn, err := grpc.DialContext(ctx, target, grpc.WithBlock(), grpc.WithInsecure(), grpc.WithBalancerName("round_robin"))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := model.NewGreeterClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	for {
		ctx, _ := context.WithTimeout(context.Background(), time.Second)
		r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("could not greet: %v", err)
		}
		log.Printf("Greeting: %s", r.Message)
		time.Sleep(time.Second * 2)
	}
}