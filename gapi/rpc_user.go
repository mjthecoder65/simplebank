package gapi

import (
	"context"

	db "github.com/mjthecoder65/simplebank/db/sqlc"
	"github.com/mjthecoder65/simplebank/pb"
	"github.com/mjthecoder65/simplebank/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	hashedPassword, err := util.HashPassword(req.GetPassword())

	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to hash password: %s", err)
	}

	arg := db.CreateUserParams{
		Username:       req.GetUsername(),
		Email:          req.GetEmail(),
		HashedPassword: hashedPassword,
		FullName:       req.GetFullName(),
	}

	user, err := server.store.CreateUser(context.Background(), arg)

	if err != nil {
		if db.ErrorCode(err) == db.UniqueViolation {
			return nil, status.Errorf(codes.AlreadyExists, "username already exist: %s", err)
		}

		return nil, status.Errorf(codes.Internal, "failed to create user: %s", err)
	}

	resp := &pb.CreateUserResponse{
		User: convertUser(user),
	}

	return resp, nil
}

func (server *Server) LoginUser(ctx context.Context, req *pb.LoginUserRequest) (*pb.LoginUserResponse, error) {
	user, err := server.store.GetUser(context.Background(), req.GetUsername())

	if err != nil {
		return nil, status.Errorf(codes.NotFound, "user with username %s was not found", req.GetUsername())
	}

	if !util.IsValidPassword(req.GetPassword(), user.HashedPassword) {
		return nil, status.Errorf(codes.InvalidArgument, "invalid password")
	}

	accessToken, err := server.maker.CreateToken(user.Username, server.config.AccessTokenDuration)

	if err != nil {
		return nil, status.Errorf(codes.Internal, "faield to created access token")
	}

	return &pb.LoginUserResponse{
		AccessToken: accessToken,
		User:        convertUser(user),
	}, nil
}
