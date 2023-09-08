package gapi

import (
	db "github.com/DMonkey83/MyFitnessApp/db/sqlc"
	"github.com/DMonkey83/MyFitnessApp/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		Email:             user.Email,
		PasswordChangedAt: timestamppb.New(user.CreatedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
