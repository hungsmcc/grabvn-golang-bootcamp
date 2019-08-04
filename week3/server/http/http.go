package http

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc"
	"week3/proto"
)

var (
	grpcPort = ":50051"
	httpPort = ":8080"
)

func StartHTTPServer() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	mux := runtime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithInsecure()}

	if err := proto.RegisterPassengerFeedbackServiceHandlerFromEndpoint(ctx, mux, grpcPort, opts); err != nil {
		log.Fatalf("failed to start HTTP gateway %v", err)
	}

	srv := http.Server{
		Addr:    httpPort,
		Handler: mux,
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			log.Println("Shutdown HTTP server...")
			if err := srv.Shutdown(context.Background()); err != nil {
				log.Fatalf("Shutdowning is failed: %v", err)
			}
		}
	}()
	srv.ListenAndServe()
}
