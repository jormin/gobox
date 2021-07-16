package token

import (
	"errors"

	"github.com/dgrijalva/jwt-go"
	"github.com/jormin/golicaidaily/library/config"
)

// Token UserTokenClaims
type UserTokenClaims struct {
	UserID   int64  `json:"user_id"`
	Mobile   string `json:"mobile"`
	FamilyID int64  `json:"family_id"`
	jwt.StandardClaims
}

// 获取 Token Key
func GetTokenKey() string {
	s, err := config.GetSecretCfg()
	if err != nil {
		return ""
	}
	return s.Token
}

// 生成Token
func GenerateToken(claims UserTokenClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(GetTokenKey()))
}

// 解析Token
func ParseToken(s string) (*UserTokenClaims, error) {
	token, err := jwt.ParseWithClaims(
		s, &UserTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte(GetTokenKey()), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*UserTokenClaims)
	if !ok || !token.Valid {
		return nil, errors.New("invalid token")
	}
	return claims, nil
}
