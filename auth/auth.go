package auth

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET"))

type JWTClaims struct {
	UsuarioId uint   `json:"usuario_id"`
	Role      string `json:"role"`
	jwt.RegisteredClaims
}

func GerarToken(UsuarioId uint, role string) (string, error) {
	claims := JWTClaims{
		UsuarioId: UsuarioId,
		Role:      role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtKey)
}

func ValidarToken(t string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(t, &JWTClaims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}

}
