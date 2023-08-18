package main

import (
	"context"
	"fmt"

	pb "gRPC/demo/gen/proto"

	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	client := pb.NewTestApiClient(conn)

	resp, err := client.Echo(context.Background(), &pb.ResponseRequest{Message: "Hello everyone!"})
	if err != nil {
		log.Println(err)
	}
	fmt.Println(resp)
	fmt.Println(resp.Message)
}
