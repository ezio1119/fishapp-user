package controllers

import (
	"context"
	"strconv"

	"github.com/ezio1119/fishapp-user/domain"
	"github.com/ezio1119/fishapp-user/pb"
	"github.com/ezio1119/fishapp-user/usecase/interactor"
	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authController struct {
	authInteractor interactor.AuthUsecase
}

func NewAuthController(a interactor.AuthUsecase) pb.UserServiceServer {
	return &authController{a}
}

func (*authController) getJwtClaimsCtx(ctx context.Context) (domain.JwtClaims, error) {
	v := ctx.Value(domain.JwtCtxKey)
	c, ok := v.(domain.JwtClaims)
	if !ok {
		return domain.JwtClaims{}, status.Errorf(codes.Internal, "failed to get jwt claims from context")
	}

	return c, nil
}

func (c *authController) CurrentUser(context.Context, *pb.CurrentUserReq) (*pb.User, error) {
	panic("a")
}

func (c *authController) UpdatePassword(context.Context, *pb.UpdatePasswordReq) (*pb.User, error) {
	panic("a")
}

func (c *authController) CreateUser(ctx context.Context, in *pb.CreateUserReq) (*pb.CreateUserRes, error) {
	// u := &domain.User{
	// 	Email:    in.Email,
	// 	Password: in.Password,
	// }

	// profileReq := &pb.CreateProfileReq{
	// 	Name:         in.ProfileName,
	// 	Introduction: in.ProfileIntroduction,
	// 	Sex:          in.ProfileSex,
	// }

	// t, p, err := c.authInteractor.CreateUser(ctx, u, profileReq)
	// if err != nil {
	// 	return nil, err
	// }

	// uProto, err := convUserProto(u)
	// if err != nil {
	// 	return nil, err
	// }

	// return &pb.CreateUserRes{
	// 	User:      uProto,
	// 	TokenPair: convTokenPairProto(t),
	// 	Profile:   p,
	// }, nil
	panic("a")
}

func (c *authController) GetUser(ctx context.Context, in *pb.GetUserReq) (*pb.User, error) {
	claims, err := c.getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}
	uID, err := strconv.ParseInt(claims.User.ID, 10, 64)
	if err != nil {
		return nil, err
	}
	u, err := c.authInteractor.GetUser(ctx, uID)
	if err != nil {
		return nil, err
	}
	return convUserProto(u)
}

func (c *authController) UpdateUser(ctx context.Context, in *pb.UpdateUserReq) (*pb.User, error) {

	// 	claims, err := c.getJwtClaimsCtx(ctx)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	uID, err := strconv.ParseInt(claims.User.ID, 10, 64)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	u := &domain.User{
	// 		ID:       uID,
	// 		Email:    in.Email,
	// 		Password: in.Password,
	// 	}
	// 	if err := c.authInteractor.UpdateUser(ctx, u, in.OldPassword); err != nil {
	// 		return nil, err
	// 	}
	// 	return convUserProto(u)
	// }

	// func (c *authController) DeleteUser(ctx context.Context, in *pb.DeleteUserReq) (*empty.Empty, error) {
	// 	claims, err := c.getJwtClaimsCtx(ctx)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	uID, err := strconv.ParseInt(claims.User.ID, 10, 64)
	// 	if err != nil {
	// 		return nil, err
	// 	}
	// 	if err := c.authInteractor.DeleteUser(ctx, uID); err != nil {
	// 		return nil, err
	// 	}
	// 	return &empty.Empty{}, nil
	panic("a")
}

func (c *authController) Login(ctx context.Context, in *pb.LoginReq) (*pb.LoginRes, error) {
	u, t, err := c.authInteractor.Login(ctx, in.Email, in.Password)
	if err != nil {
		return nil, err
	}
	uProto, err := convUserProto(u)
	if err != nil {
		return nil, err
	}
	return &pb.LoginRes{User: uProto, TokenPair: convTokenPairProto(t)}, nil
}

func (c *authController) Logout(ctx context.Context, in *pb.LogoutReq) (*empty.Empty, error) {
	claims, err := c.getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}
	if err := c.authInteractor.Logout(ctx, &claims); err != nil {
		return nil, err
	}
	return &empty.Empty{}, nil
}

func (c *authController) RefreshIDToken(ctx context.Context, in *pb.RefreshIDTokenReq) (*pb.RefreshIDTokenRes, error) {
	claims, err := c.getJwtClaimsCtx(ctx)
	if err != nil {
		return nil, err
	}
	t, err := c.authInteractor.RefreshIDToken(ctx, &claims)
	if err != nil {
		return nil, err
	}
	return &pb.RefreshIDTokenRes{TokenPair: convTokenPairProto(t)}, nil
}
