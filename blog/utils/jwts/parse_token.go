package jwts

import (
	"blog/global"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
)

// 解析
func ParseToken(tokenString string) (*CustomClaims, error) {
	MySecret = []byte(global.Config.Jwy.Secret)
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return MySecret, nil
	})

	if err != nil {
		global.Log.Error(fmt.Sprintf("token parse error: %v", err.Error()))
		return nil, err
	}
	if claim, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claim, nil
	}
	return nil, errors.New("invalid token")
}
