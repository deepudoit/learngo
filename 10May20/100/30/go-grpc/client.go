package main

import (
	"context"
	"log"

	"github.com/deepudoit/go-grpc/chat"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect", err)
	}
	defer conn.Close()
	c := chat.NewChatServiceClient(conn)

	message := chat.Message{
		Body: "I'm from client",
	}

	res, err := c.SayHello(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling SayHello", err)
	}

	log.Printf("Response from server", res.Body)
}
