package auth

import (
	"context"
	"fmt"
	"time"


	"MuchUp/backend/pkg/auth"
	"github.com/golang-jwt/jwt/v5"
)

type JWTValidator struct {
	secretKey []byte
	issuer    string
	audience  string
}

func NewJWTValidator(secretKey, issuer, audience string) *JWTValidator {
	return &JWTValidator{
		secretKey: []byte(secretKey),
		issuer:    issuer,
		audience:  audience,
	}
}

func (v *JWTValidator) ValidateToken(ctx context.Context, tokenString string) (*auth.JWTClaims, error) {
	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	default:
	}

	token, err := jwt.ParseWithClaims(tokenString, &auth.JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return v.secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to parse token: %w", err)
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	claims, ok := token.Claims.(*auth.JWTClaims)
	if !ok {
		return nil, fmt.Errorf("invalid claims type")
	}

	if v.issuer != "" && claims.Issuer != v.issuer {
		return nil, fmt.Errorf("invalid issuer")
	}

	if v.audience != "" && claims.Audience[0] != v.audience {
		return nil, fmt.Errorf("invalid audience")
	}

	return claims, nil
}

func (v *JWTValidator) GenerateToken(ctx context.Context, userID, groupID string) (string, error) {
	now := time.Now()
	claims := &auth.JWTClaims{
		UserID:  userID,
		GroupID: groupID,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    v.issuer,
			Audience:  []string{v.audience},
			ExpiresAt: jwt.NewNumericDate(now.Add(24 * time.Hour)), 
			IssuedAt:  jwt.NewNumericDate(now),
			NotBefore: jwt.NewNumericDate(now),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(v.secretKey)
}

func (v *JWTValidator) RefreshToken(ctx context.Context, tokenString string) (string, error) {
	claims, err := v.ValidateToken(ctx, tokenString)
	if err != nil {
		return "", fmt.Errorf("failed to validate token for refresh: %w", err)
	}

	return v.GenerateToken(ctx, claims.UserID, claims.GroupID)
}
