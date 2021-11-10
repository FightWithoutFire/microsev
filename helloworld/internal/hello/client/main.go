package main

import (
	"fmt"
	"google.golang.org/grpc"
	pb "helloworld/internal/conf/hello"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:5201", grpc.WithInsecure())
	if err != nil {
		return
	}
	defer conn.Close()

	client := pb.NewHelloClient(conn)
	response, err := client.SayHello(nil, nil, nil)
	fmt.Printf(response.Message)

}
