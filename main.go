package main

import (
	"context"
	"errors"
	"grpc-course-protobuf/pb/chat"
	"grpc-course-protobuf/pb/user"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// unary
type userService struct {
	user.UnimplementedUserServiceServer // placeholder untuk bantuan semua rpc yang ada (ada 4 jenis rpc)
}

// client streaming
type chatService struct {
	chat.UnimplementedChatServiceServer
}

// unary
func (us *userService) CreateUser(ctx context.Context, userReq *user.User) (*user.CreateResponse, error) {
	log.Println("user created")
	return &user.CreateResponse{
		Message: "user created",
	}, nil
}

// client streaming
func (cs *chatService) SendMessage(stream grpc.ClientStreamingServer[chat.ChatMessage, chat.ChatResponse]) error {
	for {
		req, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return status.Errorf(codes.Unknown, "error receive mesage %v", err)
		}

		log.Printf("receive message : %s to %d ", req.Content, req.UserId)
	}

	return stream.SendAndClose(&chat.ChatResponse{
		Message: "Thanks for the message!",
	})

}

// func (UnimplementedChatServiceServer) ReceiveMessage(*ReceiveMessageRequest, grpc.ServerStreamingServer[ChatMessage]) error {
// 	return status.Errorf(codes.Unimplemented, "method ReceiveMessage not implemented")
// }
// func (UnimplementedChatServiceServer) Chat(grpc.BidiStreamingServer[ChatMessage, ChatMessage]) error {
// 	return status.Errorf(codes.Unimplemented, "method Chat not implemented")
// }

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("there is error in your network listen", err)
	}

	serv := grpc.NewServer()

	user.RegisterUserServiceServer(serv, &userService{})
	chat.RegisterChatServiceServer(serv, &chatService{})

	reflection.Register(serv)

	if err := serv.Serve(listen); err != nil {
		log.Fatal("error running server", err)
	}
}
