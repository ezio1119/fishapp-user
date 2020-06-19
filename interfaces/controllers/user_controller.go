package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/ezio1119/fishapp-user/domain"
	"github.com/ezio1119/fishapp-user/pb"
	"github.com/ezio1119/fishapp-user/usecase/interactor"
	"github.com/golang/protobuf/ptypes/empty"
)

type userController struct {
	userInteractor interactor.UserUsecase
}

func NewUserController(a interactor.UserUsecase) pb.UserServiceServer {
	return &userController{a}
}

func (c *userController) CurrentUser(ctx context.Context, in *pb.CurrentUserReq) (*pb.User, error) {
	claims, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}

	uID, err := strconv.ParseInt(claims.User.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	u, err := c.userInteractor.GetUser(ctx, uID)
	if err != nil {
		return nil, err
	}

	return convUserProto(u)
}

func (c *userController) GetUser(ctx context.Context, in *pb.GetUserReq) (*pb.User, error) {
	fmt.Println("a")
	u, err := c.userInteractor.GetUser(ctx, in.Id)
	if err != nil {

		return nil, err
	}

	return convUserProto(u)
}

func (c *userController) CreateUser(ctx context.Context, in *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	pbSex, err := convSex(in.Sex)
	if err != nil {
		return nil, err
	}

	u := &domain.User{
		Email:        in.Email,
		Password:     in.Password,
		Name:         in.Name,
		Introduction: in.Introduction,
		Sex:          pbSex,
	}

	t, err := c.userInteractor.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	pbUser, err := convUserProto(u)
	if err != nil {
		return nil, err
	}

	return &pb.CreateUserRes{
		User:      pbUser,
		TokenPair: convTokenPairProto(t),
	}, nil
}

func (c *userController) UpdateUser(ctx context.Context, in *pb.UpdateUserReq) (*pb.User, error) {

	claims, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}

	uID, err := strconv.ParseInt(claims.User.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	u := &domain.User{
		ID:           uID,
		Email:        in.Email,
		Introduction: in.Introduction,
	}

	// トークンからユーザーIDを取り出しているので認可なし
	if err := c.userInteractor.UpdateUser(ctx, u); err != nil {
		return nil, err
	}
	return convUserProto(u)
}

func (c *userController) UpdatePassword(ctx context.Context, in *pb.UpdatePasswordReq) (*empty.Empty, error) {
	claims, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}

	uID, err := strconv.ParseInt(claims.User.ID, 10, 64)
	if err != nil {
		return nil, err
	}

	if err := c.userInteractor.UpdatePassword(ctx, uID, in.OldPassword, in.NewPassword); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (c *userController) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginRes, error) {
	u, t, err := c.userInteractor.Login(ctx, in.Email, in.Password)
	if err != nil {
		return nil, err
	}
	uProto, err := convUserProto(u)
	if err != nil {
		return nil, err
	}
	return &pb.LoginRes{User: uProto, TokenPair: convTokenPairProto(t)}, nil
}

func (c *userController) Logout(ctx context.Context, in *pb.LogoutReq) (*empty.Empty, error) {
	claims, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}

	if err := c.userInteractor.Logout(ctx, &claims); err != nil {
		return nil, err
	}

	return &empty.Empty{}, nil
}

func (c *userController) RefreshIDToken(ctx context.Context, in *pb.RefreshIDTokenReq) (*pb.RefreshIDTokenRes, error) {
	claims, err := getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}

	t, err := c.userInteractor.RefreshIDToken(ctx, &claims)
	if err != nil {
		return nil, err
	}

	return &pb.RefreshIDTokenRes{TokenPair: convTokenPairProto(t)}, nil
}
