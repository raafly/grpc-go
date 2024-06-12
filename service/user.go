package service

import (
	"context"

	pb "github.com/raafly/gRPC-server/pb"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"gorm.io/gorm"
)

type UserService struct {
	pb.UnimplementedUserServiceServer
	DB gorm.DB
}

func (u *UserService) GetUserData(ctx context.Context, id *pb.ID) (*pb.User, error) {
	var user pb.User

	err := u.DB.Table("users").Where("id = ?", id).Take(&user).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &user, nil
}

func (u *UserService) GetAllUser(context.Context, *pb.Empty) (*pb.Users, error) {
	var user []pb.User

	err := u.DB.Table("users").Find(&user).Error
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Konversi slice []pb.User menjadi slice []*pb.User
	userPointers := make([]*pb.User, len(user))

	// Membungkus hasil query dalam struct pb.Users dan mengembalikannya sebagai pointer
	return &pb.Users{Users: userPointers}, nil
}
