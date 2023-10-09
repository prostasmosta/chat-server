package main

import (
	"context"
	"log"
	"time"

	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	grpcChat "github.com/prostasmosta/chat-server/grpc/pkg/chat_v1"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect ot server: %v", err)
	}
	defer conn.Close()

	c := grpcChat.NewChatV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	userNames := []string{"Peter", "Sam", "John"}
	
	r, err := c.Create(ctx, &grpcChat.CreateRequest{Usernames: userNames})
	if err != nil {
		log.Fatalf("failed to create new chat: %v", err)
	}

	log.Printf(color.RedString("Chat info:\n"), color.GreenString("%+v", r.GetId()))
}
