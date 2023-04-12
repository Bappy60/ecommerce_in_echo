package tokens

import (
	"time"

	"github.com/Bappy60/ecommerce_in_echo/pkg/config"
	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Email string
	UserId uint64
	jwt.StandardClaims
}


func TokenGenerator(email string,userID uint64) (signedtoken string, err error) {
	claims := &SignedDetails{
		Email: email,
		UserId : userID,
		StandardClaims: jwt.StandardClaims{
			IssuedAt:  time.Now().Local().Unix(),
			ExpiresAt: time.Now().Local().Add(time.Hour * time.Duration(24)).Unix(),
		},
	}
	token, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(config.LocalConfig.SECRET_KEY))
	if err != nil {
		return "", err
	}
	return token, err
}

func ValidateToken(signedtoken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(signedtoken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.LocalConfig.SECRET_KEY), nil
	})

	if err != nil {
		msg = err.Error()
		return
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "The Token is invalid"
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "token is expired"
		return
	}
	return claims, msg
}