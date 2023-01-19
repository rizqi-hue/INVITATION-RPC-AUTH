package utils

import (
	"errors"
	"time"

	"github.com/INVITATION-RPC-AUTH/domain/models"
	"github.com/golang-jwt/jwt/v4"
)

type JwtWrapper struct {
	SecretKey       string
	Issuer          string
	ExpirationHours int64
}

type jwtClaims struct {
	jwt.RegisteredClaims
	Id    int64
	Email string
}

func (w *JwtWrapper) GenerateAccessToken(user models.User) (signedToken string, err error) {
	// expirationTime := time.Now().Add(5 * time.Minute)

	expirationTime := time.Now().Local().Add(time.Hour * time.Duration(w.ExpirationHours)).Unix()

	claims := &jwtClaims{
		Id:    user.Id,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Unix(expirationTime, 0)),
			Issuer:    w.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err = token.SignedString([]byte(w.SecretKey))

	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func (w *JwtWrapper) ValidateToken(signedToken string) (claims *jwtClaims, err error) {

	token, err := jwt.ParseWithClaims(
		signedToken,
		&jwtClaims{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(w.SecretKey), nil
		},
	)

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("err signature invalid")
		}

		return nil, errors.New("err bad request")
	}

	claims, ok := token.Claims.(*jwtClaims)

	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	// if claims.ExpiresAt < jwt.NewNumericDate(time.Now()) {
	// 	return nil, errors.New("JWT is expired")
	// }

	return claims, nil
}
