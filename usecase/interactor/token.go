package interactor

import (
	"crypto/ecdsa"
	"log"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ezio1119/fishapp-user/conf"
	"github.com/ezio1119/fishapp-user/domain"
	"github.com/google/uuid"
)

func genTokenPair(id string) (*domain.TokenPair, error) {
	idToken, err := genToken(id, "id_token", conf.C.Auth.IDTokenExpSec)
	if err != nil {
		return nil, err
	}
	refreshToken, err := genToken(id, "refresh_token", conf.C.Auth.RtExpSec)
	if err != nil {
		return nil, err
	}
	return &domain.TokenPair{
		IDToken:      idToken,
		RefreshToken: refreshToken,
	}, nil
}

func genToken(userID string, sub string, expSec int64) (string, error) {
	expTime := time.Now().Add(time.Duration(expSec) * time.Second)
	claims := &domain.JwtClaims{
		struct {
			ID string `json:"id"`
		}{
			ID: userID,
		},
		jwt.StandardClaims{
			Id:        uuid.New().String(),
			Subject:   sub,
			ExpiresAt: expTime.Unix(),
		},
	}
	t := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	return t.SignedString(privateKey)
}

var privateKey *ecdsa.PrivateKey

func init() {
	var err error
	privateKey, err = jwt.ParseECPrivateKeyFromPEM([]byte(conf.C.Auth.PvtJwtkey))
	if err != nil {
		log.Fatal(err)
	}
}
