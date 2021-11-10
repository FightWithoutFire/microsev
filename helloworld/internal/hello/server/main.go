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
	Address = "127.0.0.1:52100"
)

type helloService struct {
	pb.HelloServer
}

func (s helloService) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	resp := new(pb.HelloResponse)
	resp.Message = fmt.Sprintf("Hello %s.", in.Name)
	fmt.Printf("response: %s", resp.Message)
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
	fmt.Printf("server started")
	err = server.Serve(listen)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
}
