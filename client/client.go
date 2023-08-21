package main

import (
	"context"
	"fmt"

	pb "demo/gen/proto"

	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	client := pb.NewTestAPIClient(conn)

	resp, err := client.Echo(context.Background(), &pb.ResponseRequest{Message: "Hello everyone!"})
	if err != nil {
		log.Println(err)
	}
	respClient, err := client.GetUser(context.Background(), &pb.UserResponse{Name: "Rk", Age: 25})
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Echo :", resp.Message)
	fmt.Println("Client :", respClient)
}
