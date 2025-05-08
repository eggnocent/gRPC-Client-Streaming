package main

import (
	"context"
	"grpc-course-protobuf/pb/user"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type userService struct {
	user.UnimplementedUserServiceServer // placeholder untuk bantuan semua rpc yang ada (ada 4 jenis rpc)
}

func (us *userService) CreateUser(ctx context.Context, userReq *user.User) (*user.CreateResponse, error) {
	log.Println("user created")
	return &user.CreateResponse{
		Message: "user created",
	}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("there is error in your network listen", err)
	}

	serv := grpc.NewServer()

	user.RegisterUserServiceServer(serv, &userService{})

	reflection.Register(serv)

	if err := serv.Serve(listen); err != nil {
		log.Fatal("error running server", err)
	}
}
