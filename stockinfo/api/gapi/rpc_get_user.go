package gapi

import (
	"context"
	"database/sql"

	"github.com/RoyceAzure/go-stockinfo/shared/pb"
	"github.com/RoyceAzure/go-stockinfo/shared/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	payload, err := server.authorizUser(ctx)
	if err != nil {
		return nil, util.UnauthticatedError(err)
	}

	user, err := server.store.GetUser(ctx, payload.UserId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, status.Errorf(codes.NotFound, "user not found")
		}
		return nil, status.Errorf(codes.Internal, "failed to Update user : %s", err)
	}
	res := &pb.GetUserResponse{
		User: convertUser(user),
	}
	return res, nil
}
