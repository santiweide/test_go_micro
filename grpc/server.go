package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
	"net"
	"test_go_micro/grpc/consul"
	"test_go_micro/grpc/model"
)

const (
	port = ":50051"
)

type Greeter struct{
	model.UnimplementedGreeterServer
}


func (g *Greeter) TestStruct(ctx context.Context, req *model.StructRequest) (*model.StructResponse, error) {
	resp := &model.StructResponse{
		Id:         req.Id,
		KvMap:      req.KvMap,
		StringList: req.StringList,
	}
	return resp, nil
}

func (g *Greeter) TestString(ctx context.Context, req *model.StringRequest) (*model.StringResponse, error) {
	resp := &model.StringResponse{Message: req.Message}
	return resp, nil

}

func RegisterToConsul() {
	consul.RegitserService("127.0.0.1:8500", &consul.ConsulService{
		Name: "test_grpc",
		Tag:  []string{"test","grpc"},
		IP:   "127.0.0.1",
		Port: 50051,
	})
}

//health
type HealthImpl struct{}

// Check 实现健康检查接口，这里直接返回健康状态，这里也可以有更复杂的健康检查策略，比如根据服务器负载来返回
func (h *HealthImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	fmt.Print("health checking\n")
	return &grpc_health_v1.HealthCheckResponse{
		Status: grpc_health_v1.HealthCheckResponse_SERVING,
	}, nil
}

func (h *HealthImpl) Watch(req *grpc_health_v1.HealthCheckRequest, w grpc_health_v1.Health_WatchServer) error {
	return nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	model.RegisterGreeterServer(s, &Greeter{})
	grpc_health_v1.RegisterHealthServer(s, &HealthImpl{})
	RegisterToConsul()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
