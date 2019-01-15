package authentication

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
)

func RequireTokenAuthentication(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	authBackend := InitJWTAuthenticationBackend()

	token, err := request.ParseFromRequest(r, request.OAuth2Extractor, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["arg"])
		} else {
			return authBackend.PublicKey, nil
		}
	})

	if err == nil && token.Valid && !authBackend.IsInBlackList(r.Header.Get("Authorization")) {
		next(w, r)
	} else {
		w.WriteHeader(http.StatusUnauthorized)
	}
}
