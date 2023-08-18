package main

import (
	"context"
	pb "demo/gen/proto"
	"net"

	"log"

	"google.golang.org/grpc"
)

type testAPIServer struct {
	pb.UnimplementedTestAPIServer
}

func (s *testAPIServer) GetUser(ctx context.Context, req *pb.UserRequest) (*pb.UserResponse, error) {
	return &pb.UserResponse{}, nil
}
func (s *testAPIServer) Echo(ctx context.Context, req *pb.ResponseRequest) (*pb.ResponseRequest, error) {
	return req, nil
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTestAPIServer(grpcServer, &testAPIServer{})

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Println(err)
	}
}
