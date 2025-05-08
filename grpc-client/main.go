package main

import (
	"context"
	"grpc-course-protobuf/pb/user"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	clientConn, err := grpc.NewClient("localhost:8080", grpc.WithTransportCredentials(insecure.NewCredentials())) // tempat untuk memsaukan ssl dan disini di bypass dengan insecure
	if err != nil {
		log.Fatal("failed to create client", err)
	}

	userClient := user.NewUserServiceClient(clientConn)

	resp, err := userClient.CreateUser(context.Background(), &user.User{
		Id:      1,
		Age:     15,
		Balance: 12000,
		Address: &user.Address{
			Id:          "12",
			FullAddress: "Jln.Tempel",
			Province:    "DIY",
			City:        "Sleman",
		},
	})

	if err != nil {
		log.Fatal("error calling user client", err)
	}

	log.Println("Got message from server: ", resp.Message)
}
