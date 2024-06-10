package main

import (
	"fmt"

	pb "github.com/raafly/gRPC-server/pb"
	"google.golang.org/protobuf/proto"
)

func main() {
	users := &pb.Users{
		Users: []*pb.User{
			{
				Id: 1,
				Name: "rafly",
				Email: "rafliexecutor375@gmail.com",
			},
			{
				Id: 2,
				Name: "xShavaor",
				Email: "anonymousx@gmail.com",
			},
		},
	}

	result, _ := proto.Marshal(users)
	fmt.Println(result)
}