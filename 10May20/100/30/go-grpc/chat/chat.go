package chat

import (
	"context"
	"log"
)

type Server struct {
}

func (s *Server) SayHello(ctx context.Context, message *Message) (*Message, error) {
	log.Printf("Server server >>> %s\n", message.Body)
	return &Message{Body: "Yo bro, you are on Fire...."}, nil
}
