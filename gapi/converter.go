package gapi

import (
	db "github.com/mjthecoder65/simplebank/db/sqlc"
	"github.com/mjthecoder65/simplebank/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		Username:          user.Username,
		Email:             user.Email,
		FullName:          user.FullName,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CreatedAt:         timestamppb.New(user.CreatedAt),
	}
}
