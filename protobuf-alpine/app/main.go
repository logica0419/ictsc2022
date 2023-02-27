package main

import (
	"context"
	"log"
	"net"

	"app/protobuf"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type PingServer struct {
	protobuf.UnimplementedPingServiceServer
}

func (p *PingServer) Ping(context.Context, *protobuf.PingRequest) (*protobuf.PingResponse, error) {
	return &protobuf.PingResponse{Message: "Hello, ICTSC2022 Contestant!"}, nil
}

func main() {
	ping := &PingServer{}

	server := grpc.NewServer()
	protobuf.RegisterPingServiceServer(server, ping)

	reflection.Register(server)

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Panic(err)
	}

	log.Panic(server.Serve(listener))
}
