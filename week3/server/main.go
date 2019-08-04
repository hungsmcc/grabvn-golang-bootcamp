package main

import (
	"log"

	"week3/server/grpc"
	"week3/server/http"
)

func main() {
	go func() {
		log.Println("Starting  gRPC server...")
		grpc.StartGRPCServer()
	}()

	log.Println("Starting  HTTP server...")
	http.StartHTTPServer()
}
