package interactor

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/ezio1119/fishapp-user/domain"
	"github.com/ezio1119/fishapp-user/usecase/repository"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type authInteractor struct {
	userRepository      repository.UserRepository
	blackListRepository repository.BlackListRepository
	ctxTimeout          time.Duration
}

func NewAuthInteractor(
	u repository.UserRepository,
	b repository.BlackListRepository,
	t time.Duration,
) AuthUsecase {
	return &authInteractor{u, b, t}
}

type AuthUsecase interface {
	// CreateUser(ctx context.Context, u *domain.User, profileReq *pb.CreateProfileReq) (*domain.TokenPair, *pb.Profile, error)
	GetUser(ctx context.Context, id int64) (*domain.User, error)
	UpdateUser(ctx context.Context, u *domain.User, oldPass string) error
	DeleteUser(ctx context.Context, id int64) error
	Login(ctx context.Context, email string, pass string) (*domain.User, *domain.TokenPair, error)
	Logout(ctx context.Context, jwtClaims *domain.JwtClaims) error
	RefreshIDToken(ctx context.Context, jwtClaims *domain.JwtClaims) (*domain.TokenPair, error)
}

func (i *authInteractor) GetUser(ctx context.Context, id int64) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()

	return i.userRepository.GetUser(ctx, id)
}

// func (i *authInteractor) CreateUser(ctx context.Context, u *domain.User, profileReq *pb.CreateProfileReq) (*domain.TokenPair, *pb.Profile, error) {
// 	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
// 	defer cancel()

// 	encryptedPass, err := genEncryptedPass(u.Password)
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	u.EncryptedPassword = encryptedPass

// 	if err := i.userRepository.CreateUser(ctx, u); err != nil {
// 		return nil, nil, err

// 	}
// 	profileReq.UserId = u.ID

// 	// CreateProfileに失敗したらDeleteUserする
// 	p, err := i.profileRepository.CreateProfile(ctx, profileReq)
// 	if err != nil {
// 		if err := i.userRepository.DeleteUser(ctx, u.ID); err != nil {
// 			return nil, nil, err
// 		}
// 		return nil, nil, err
// 	}

// 	t, err := genTokenPair(strconv.FormatInt(u.ID, 10))
// 	if err != nil {
// 		return nil, nil, err
// 	}

// 	return t, p, nil
// }

func (i *authInteractor) UpdateUser(ctx context.Context, u *domain.User, oldPass string) error {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	res, err := i.userRepository.GetUser(ctx, u.ID)
	if err != nil {
		return err
	}
	if err := compareHashAndPass(res.EncryptedPassword, oldPass); err != nil {
		return status.Errorf(codes.Unauthenticated, "wrong password: %s", err)
	}
	encryptedPass, err := genEncryptedPass(u.Password)
	if err != nil {
		return err
	}
	u.EncryptedPassword = encryptedPass
	if err := i.userRepository.UpdateUser(ctx, u); err != nil {
		return err
	}
	u.CreatedAt = res.CreatedAt
	fmt.Printf("%#v", u)
	return nil
}

func (i *authInteractor) DeleteUser(ctx context.Context, id int64) error {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	return i.userRepository.DeleteUser(ctx, id)
}

func (i *authInteractor) Login(ctx context.Context, email string, pass string) (*domain.User, *domain.TokenPair, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	u, err := i.userRepository.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, nil, err
	}
	if err := compareHashAndPass(u.EncryptedPassword, pass); err != nil {
		return nil, nil, status.Errorf(codes.Unauthenticated, "token is blacklisted")
	}
	tokenPair, err := genTokenPair(strconv.FormatInt(u.ID, 10))
	if err != nil {
		return nil, nil, err
	}
	return u, tokenPair, nil
}

func (i *authInteractor) Logout(ctx context.Context, c *domain.JwtClaims) error {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	exp := time.Unix(c.ExpiresAt, 0).Sub(time.Now())
	if _, err := i.blackListRepository.SetNX(c.Id, exp); err != nil {
		return err
	}
	return nil
}

func (i *authInteractor) RefreshIDToken(ctx context.Context, c *domain.JwtClaims) (*domain.TokenPair, error) {
	ctx, cancel := context.WithTimeout(ctx, i.ctxTimeout)
	defer cancel()
	res, err := i.blackListRepository.Exists(c.Id)
	if err != nil {
		return nil, err
	}
	if res == 1 {
		return nil, status.Error(codes.Unauthenticated, "token is blacklisted")
	}

	tokenPair, err := genTokenPair(c.User.ID)
	if err != nil {
		return nil, err
	}
	exp := time.Unix(c.ExpiresAt, 0).Sub(time.Now()) // トークンの有効期限 - 現在の時間
	if _, err := i.blackListRepository.SetNX(c.Id, exp); err != nil {
		return nil, err
	}
	return tokenPair, nil
}
