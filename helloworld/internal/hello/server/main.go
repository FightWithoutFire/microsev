package main

import (
	"fmt"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "helloworld/internal/conf/hello"
	"net"
)

const (
	Address = "127.0.0.1:5210"
)

type helloService struct {
}

func (s helloService) mustEmbedUnimplementedHelloServer() {
	panic("implement me")
}

func (s helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error)  {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)

	return resp, nil
}

var HelloService = helloService{}
func main() {

	listen, err := net.Listen("tcp", Address)
	if err != nil {
		grpclog.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterHelloServer(server, HelloService)
	grpclog.Println("Listen on " + Address)
	server.Serve(listen)
}
