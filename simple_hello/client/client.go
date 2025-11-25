package main

import (
	"context"
	"log"
	simple "simple_hello/proto/generated"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Not able to Conntect localhost:50051 %v", err)
	}
	defer conn.Close()
	client := simple.NewSimpleServiceClient(conn)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	response, err := client.SendMessage(
		ctx, &simple.SimpleRequest{
			MyMessage: "Banti bol rah bhai client side se",
		})

	if err != nil {
		log.Fatalf("Not able to send the message :%v", err)
	}

	log.Println("\nGot the reply bro \n", response, "\n")
}
