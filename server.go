package main

import (
	"context"
	pb "demo/gen/proto"
	"fmt"
	"net"
	"strings"

	"log"

	"google.golang.org/grpc"
)

type testAPIServer struct {
	pb.UnimplementedTestAPIServer
}

func (s *testAPIServer) GetUser(ctx context.Context, req *pb.UserResponse) (*pb.UserResponse, error) {
	//fetch email based on name and age from db
	req.Email = "rk@gridinfocom.com"
	return req, nil
}
func (s *testAPIServer) Echo(ctx context.Context, req *pb.ResponseRequest) (*pb.ResponseRequest, error) {
	req.Message = strings.Replace(req.Message, "!", ". Greetings from server side!", 1)
	return req, nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Server Started")
	grpcServer := grpc.NewServer()
	pb.RegisterTestAPIServer(grpcServer, &testAPIServer{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Println(err)
	}
}
