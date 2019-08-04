package main

import (
	"context"
	"net"

	"google.golang.org/grpc"
	// "google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	// "google.golang.org/grpc/status"
	//pb "grab.com/demo/pb"
	pb "grab.com/demo/pb"
)

type greeter struct{}

func (s *greeter) SayHello(context.Context, *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello back"}, nil

	// TODO: return error
	// return nil, status.Error(codes.NotFound, "I don't know who you are")
}

func main() {
	// create a server struct
	server := grpc.NewServer()
	// register handlers
	pb.RegisterGreeterServer(server, &greeter{})
	// create the network socket for the server to listen to
	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}
	// reflection enables us to use grpc_cli for local development and debugging
	reflection.Register(server)
	// servers starts to accept requests
	server.Serve(listener)
}
