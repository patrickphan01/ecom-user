package main

import (
	"log"
	"net"

	config "github.com/phankieuphu/ecom-user/config"
	pb "github.com/phankieuphu/ecom-user/services"
	"google.golang.org/grpc"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error load env", err)
	}

	config.GetDb()

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpServer := grpc.NewServer()
	pb.RegisterUserServer(grpServer, &pb.UserService{})

	log.Println("gRPC server listening on port 50051...")
	if err := grpServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
