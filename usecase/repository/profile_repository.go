package repository

import (
	"context"

	"github.com/ezio1119/fishapp-auth/pb"
)

type ProfileRepository interface {
	CreateProfile(ctx context.Context, req *pb.CreateProfileReq) (*pb.Profile, error)
}
