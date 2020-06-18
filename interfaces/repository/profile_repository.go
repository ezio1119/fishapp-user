package repository

import (
	"context"

	"github.com/ezio1119/fishapp-auth/pb"
	"github.com/ezio1119/fishapp-auth/usecase/repository"
)

type profileRepository struct {
	client pb.ProfileServiceClient
}

func NewProfileRepository(c pb.ProfileServiceClient) repository.ProfileRepository {
	return &profileRepository{c}
}

func (r *profileRepository) CreateProfile(ctx context.Context, req *pb.CreateProfileReq) (*pb.Profile, error) {
	return r.client.CreateProfile(ctx, req)
}
