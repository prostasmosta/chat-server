package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/emptypb"
	
	grpcChat "github.com/prostasmosta/chat-server/grpc/pkg/chat_v1"
)

const grpcPort = 50051

type server struct {
	grpcChat.UnimplementedChatV1Server
}

func (s *server) Create(ctx context.Context, req *grpcChat.CreateRequest) (*grpcChat.CreateResponse, error) {
	fakeId := gofakeit.Int64()
	log.Printf("Chat id: %d", fakeId)

	return &grpcChat.CreateResponse{
		Id: fakeId,
	}, nil
}

func (s *server) Delete(ctx context.Context, req *grpcChat.DeleteRequest) (*emptypb.Empty, error) {
	log.Printf("Chat id: %d", req.GetId())

	return &emptypb.Empty{}, nil
}

func (s *server) SendMessage(ctx context.Context, req *grpcChat.SendMessageRequest) (*emptypb.Empty, error) {
	log.Printf("Chat tmstmp: %v", req.GetTimestamp())

	return &emptypb.Empty{}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	grpcChat.RegisterChatV1Server(s, &server{})

	log.Printf("server listening at: %v", lis.Addr())

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
