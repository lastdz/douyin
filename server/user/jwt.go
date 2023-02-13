package user

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
)

type JwtService struct {
	secret []byte
}

func NewJwtService(secret []byte) *JwtService {
	return &JwtService{secret: secret}
}

// GenerateToken 根据 Username 生成一个 Token
func (js *JwtService) GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"usr": username,
	})

	tokenString, err := token.SignedString(js.secret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// ValidateToken 校验 Token 是否有效
func (js *JwtService) ValidateToken(tokenStr string, username string) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return js.secret, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if !ok {
			return errors.New("couldn't handle this Token")
		}
		if claims["usr"] != username {
			return errors.New("validation failed")
		}
	} else if errors.Is(err, jwt.ErrTokenMalformed) {
		return errors.New("wrong Token format" + err.Error())
	} else if errors.Is(err, jwt.ErrTokenExpired) || errors.Is(err, jwt.ErrTokenNotValidYet) {
		// Token is either expired or not active yet
		return errors.New("timing is everything:" + err.Error())
	} else {
		return errors.New("Couldn't handle this Token:" + err.Error())
	}
	return nil
}
