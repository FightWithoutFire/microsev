package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "helloworld/internal/conf/hello"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:52100", grpc.WithInsecure())
	if err != nil {
		grpclog.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	req := pb.HelloRequest{Name: "good"}
	c := pb.NewHelloClient(conn)
	response, err := c.SayHello(context.Background(), &req)
	if err != nil {
		grpclog.Fatalf("Failed to get response: %v", err)
	}
	fmt.Printf(response.Message)

}
