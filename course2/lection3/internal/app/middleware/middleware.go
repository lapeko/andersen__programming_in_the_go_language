package middleware

import (
	jwtmiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/form3tech-oss/jwt-go"
)

var (
	SecretToken []byte      = []byte("MegaSecretToken ( ! )")
	keyFunc     jwt.Keyfunc = func(_ *jwt.Token) (interface{}, error) {
		// Our token must be signed using this data.
		return SecretToken, nil
	}
)

var JWTMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
	ValidationKeyGetter: keyFunc,
	SigningMethod:       jwt.SigningMethodHS256,
})
