package gapi

import (
	"github.com/RoyceAzure/go-stockinfo-api/pb"
	db "github.com/RoyceAzure/go-stockinfo-project/db/sqlc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(user db.User) *pb.User {
	return &pb.User{
		UserName:          user.UserName,
		Email:             user.Email,
		SsoIdentifer:      user.SsoIdentifer.String,
		PasswordChangedAt: timestamppb.New(user.PasswordChangedAt),
		CrDate:            timestamppb.New(user.CrDate),
	}
}
