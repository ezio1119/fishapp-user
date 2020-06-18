package controllers

import (
	"github.com/ezio1119/fishapp-user/domain"
	"github.com/ezio1119/fishapp-user/pb"
	"github.com/golang/protobuf/ptypes"
)

func convUserProto(u *domain.User) (*pb.User, error) {
	updatedAt, err := ptypes.TimestampProto(u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	createdAt, err := ptypes.TimestampProto(u.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &pb.User{
		Id:        u.ID,
		Email:     u.Email,
		UpdatedAt: updatedAt,
		CreatedAt: createdAt,
	}, nil
}

func convTokenPairProto(tp *domain.TokenPair) *pb.TokenPair {
	return &pb.TokenPair{
		IdToken:      tp.IDToken,
		RefreshToken: tp.RefreshToken,
	}
}
