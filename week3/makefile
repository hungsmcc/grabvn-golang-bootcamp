VERSION := $(shell git describe --tags)

#Check go
CC = go
go = $(shell which go 2> /dev/null)
ifeq (, $(go))
    @printf "\e[91mGo not found!"
endif

#gen proto
generate:
	protoc -I proto/ --proto_path=$(HOME)/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis proto/feedback.proto --go_out=plugins=grpc:proto
	protoc -I proto/ --proto_path=$(HOME)/go/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:proto proto/feedback.proto
	echo "Generate done"

#run app
start-server:
	go run ./server/main.go

start-client:
	go run ./client/main.go