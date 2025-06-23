package middleware

import (
	"context"
	"net/http"
	"strings"

	"MuchUp/backend/pkg/auth"
)

type contextKey string

const (
    UserIDContextKey  contextKey = "userID"
    GroupIDContextKey contextKey = "groupID"
    ClaimsContextKey  contextKey = "claims"
)


func JWTMiddleware(next http.Handler,validator auth.TokenValidator ) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var ctx context.Context
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing authorization header", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "invalid token format", http.StatusUnauthorized)
			return
		}
		claims,err := validator.ValidateToken(ctx,tokenString)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		
		ctx = context.WithValue(r.Context(),UserIDContextKey,claims.UserID)
	    ctx = context.WithValue(ctx,GroupIDContextKey,claims.GroupID)
		ctx = context.WithValue(ctx,ClaimsContextKey,claims)


		next.ServeHTTP(w,r.WithContext(ctx))
	})
}

