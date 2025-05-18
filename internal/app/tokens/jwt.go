package tokens

import (
	"fmt"
	"github.com/go-chi/jwtauth/v5"
	"github.com/golang-jwt/jwt/v4"
	jwt2 "github.com/lestrrat-go/jwx/v2/jwt"
	"time"
)

type JWTWorker struct {
	*jwtauth.JWTAuth
	secret string
	exp    time.Duration
}

func NewJWTWorker(secret string, exp time.Duration) *JWTWorker {
	auth := jwtauth.New(
		"HS512",
		[]byte(secret),
		nil,
		jwt2.WithAcceptableSkew(exp),
	)

	return &JWTWorker{
		JWTAuth: auth,
		secret:  secret,
		exp:     exp,
	}
}

func (j *JWTWorker) Encode(claims map[string]any) (string, error) {
	claims["iss"] = "github.com/StratuStore/auth"
	claims["exp"] = jwt.NewNumericDate(time.Now().Add(j.exp))
	claims["iat"] = jwt.NewNumericDate(time.Now())

	t := jwt.NewWithClaims(jwt.SigningMethodHS512, jwt.MapClaims(claims))

	token, err := t.SignedString([]byte(j.secret))
	if err != nil {
		return "", fmt.Errorf("unable to sign token: %w", err)
	}

	return token, nil
}
