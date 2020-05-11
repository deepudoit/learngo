package main

import (
	"github.com/deepudoit/go-grpc/chat"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Failed to start TCP conn: %v", err)
	}

	s := chat.Server{}
	gSer := grpc.NewServer()
	chat.RegisterChatServiceServer(gSer, &s)

	if err := gSer.Serve(lis); err != nil {
		log.Fatalf("Failed to start gRPC server: %v", err)
	}

}
