package tokens

import (
	"time"

	"github.com/Bappy60/ecommerce_in_echo/pkg/config"
	jwt "github.com/dgrijalva/jwt-go"
)

type SignedDetails struct {
	Email string
	jwt.StandardClaims
}


func TokenGenerator(email string) (signedtoken string, err error) {
	claims := &SignedDetails{
		Email: email,
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
