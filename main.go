package main

import (
	"log"
	"net"

	"github.com/raafly/gRPC-server/db"
	pb "github.com/raafly/gRPC-server/pb"
	"github.com/raafly/gRPC-server/service"
	"google.golang.org/grpc"
)

func main() {
	db := db.NewDB()

	netListen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("cannot listen port %w", err.Error())
	}

	gprcServer := grpc.NewServer()
	portService := service.UserService{DB: *db}
	pb.RegisterUserServiceServer(gprcServer, &portService)


	log.Println("server running...")
	if err := gprcServer.Serve(netListen); err != nil {
		log.Println("failed to serve %w", err.Error())
	}
}