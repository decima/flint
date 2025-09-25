package security

import (
	"flint/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct {
	jwtSecret []byte
}

func NewJwt(config *config.Config) *Jwt {
	return &Jwt{
		jwtSecret: []byte(config.Security.JWTSecret),
	}
}
func (j *Jwt) GenerateToken(user string, role Role) (string, error) {
	tokenLifetime := time.Hour * 2
	claims := jwt.MapClaims{}
	claims["user"] = user
	claims["role"] = role
	claims["exp"] = time.Now().Add(tokenLifetime).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := j.jwtSecret // Replace with your secret key
	return token.SignedString(secret)
}

func (j *Jwt) GenerateRefreshToken(user string) (string, error) {
	tokenLifetime := time.Hour * 24 * 7 // 7 days
	claims := jwt.MapClaims{}
	claims["user"] = user
	claims["refresh"] = true
	claims["exp"] = time.Now().Add(tokenLifetime).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	secret := j.jwtSecret // Replace with your secret key
	return token.SignedString(secret)
}

func (j *Jwt) ValidateToken(tokenString string) (string, Role, error) {
	secret := j.jwtSecret // Replace with your secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secret, nil
	})
	if err != nil {
		return "", Anon, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := claims["user"].(string)
		role := RoleFromString(claims["role"].(string))
		return user, role, nil
	}
	return "", Anon, jwt.ErrInvalidKey
}

func (j *Jwt) ValidateRefreshToken(tokenString string) (string, error) {
	secret := j.jwtSecret // Replace with your secret key
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return secret, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := claims["user"].(string)
		return user, nil
	}
	return "", jwt.ErrInvalidKey
}
