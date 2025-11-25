package main

import (
	"context"
	"log"
	"net"
	simple "simple_hello/proto/generated"

	"google.golang.org/grpc"
)

type UnInplementServer struct {
	simple.UnimplementedSimpleServiceServer
}

func (s *UnInplementServer) SendMessage(ctx context.Context, req *simple.SimpleRequest) (*simple.SimpleResponse, error) {
	log.Println("Received message from client as req :", req.MyMessage)
	return &simple.SimpleResponse{
		Reply:    req.MyMessage,
		Received: true,
	}, nil
}

func main() {
	listening, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to liten...%v", err)
	}
	server := grpc.NewServer()
	simple.RegisterSimpleServiceServer(server, &UnInplementServer{})
	log.Println("Server is listening on 50051")

	if err := server.Serve(listening); err != nil {
		log.Fatalf("Failed to serve :%v", err)
	}

}
