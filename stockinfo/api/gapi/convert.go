package gapi

import (
	db "github.com/RoyceAzure/go-stockinfo/repository/db/sqlc"
	"github.com/RoyceAzure/go-stockinfo/shared/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		UserId:            user.UserID,
		UserName:          user.UserName,
		Email:             user.Email,
		SsoIdentifer:      user.SsoIdentifer.String,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CrDate:            timestamppb.New(user.CrDate),
	}
}
