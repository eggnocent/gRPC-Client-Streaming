package main

import (
	"context"
	"grpc-course-protobuf/pb/chat"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	clientConn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials())) // tempat untuk memsaukan ssl dan disini di bypass dengan insecure
	if err != nil {
		log.Fatal("failed to create client", err)
	}

	chatClient := chat.NewChatServiceClient(clientConn)
	stream, err := chatClient.SendMessage(context.Background())
	if err != nil {
		log.Fatal("failed to send message: ", err)
	}

	err = stream.Send(&chat.ChatMessage{
		UserId:  1,
		Content: "hello world",
	})

	if err != nil {
		log.Fatal("failed to send via stream", err)
	}

	err = stream.Send(&chat.ChatMessage{
		UserId:  1,
		Content: "hello again",
	})

	if err != nil {
		log.Fatal("failed to send via stream", err)
	}
	time.Sleep(5 * time.Second)

	err = stream.Send(&chat.ChatMessage{
		UserId:  1,
		Content: "hello there!",
	})

	if err != nil {
		log.Fatal("failed to send via stream", err)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal("failed to close", err)
	}
	log.Println("connection is close", res.Message)
}
