package main

import (
	"context"
	"fmt"
	"os"
	"time"

	// "time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "grab.com/demo/pb"
)

func main() {
	// specify server address, and no authentication
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("Failed to connect", err)
		os.Exit(1)
	}
	// close the connection when the function finishes
	defer func() {
		_ = conn.Close()
	}()

	// create a client using the generated code
	c := pb.NewGreeterClient(conn)
	// ctx := context.Background()

	// TODO: Add timeout
	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()

	// call SayHello function with a parameter name
	resp, err := c.SayHello(ctx, &pb.HelloRequest{Name: "name"})

	// if err != nil {
	// 	fmt.Println("Failed to call server", err)
	// 	os.Exit(1)
	// }

	// TODO: Handle error
	if err != nil {
		grpcErr, ok := status.FromError(err)
		if ok {
			switch grpcErr.Code() {
			case codes.NotFound:
				fmt.Println("Not Found: ", grpcErr.Message())
			default:
				fmt.Println("Unexpected error: ", grpcErr.Code())
			}
		} else {
			fmt.Println("Failed to call server", err)
		}
		return
	}

	fmt.Println("Server reply:", resp.Message)
}
