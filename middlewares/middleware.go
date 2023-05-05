package middlewares

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"test-dikurium/Token"
)

type AuthMiddlewares struct {
	TokenSrv Token.JWTMaker
}

func NewAuthMiddlewares(tokenSrv Token.JWTMaker) *AuthMiddlewares {
	return &AuthMiddlewares{TokenSrv: tokenSrv}
}

func (a AuthMiddlewares) Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header from the request

		tokenHeader := r.Header.Get("Authorization")
		if tokenHeader == "" {
			next.ServeHTTP(w, r)
			return
		}

		fields := strings.Fields(tokenHeader)
		if len(fields) < 3 {
			next.ServeHTTP(w, r)
			return
		}

		AuthorizationType := strings.ToLower(fields[0])
		if AuthorizationType != "bearer" {
			http.Error(w, "invalid auth format", http.StatusUnauthorized)
			return
		}

		accessToken := fields[2]
		data, err := a.TokenSrv.VerifyToken(accessToken)
		if err != nil {
			http.Error(w, fmt.Sprintf("token not verified: %v", err), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "email", data.Email)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
