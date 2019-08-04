package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	pb "grab.com/demo/pb"
)

var (
	client pb.GreeterClient
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
	client = pb.NewGreeterClient(conn)
	ctx := context.Background()

	ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()

	router := gin.Default()
	router.GET("/hello", sayHello)
	router.Run(":8088")
}

func sayHello(g *gin.Context) {
	// call SayHello function with a parameter name
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "name"})

	if err != nil {
		g.String(500, "failed to call server")
		return
	}

	g.JSON(200, resp.Message)
}
