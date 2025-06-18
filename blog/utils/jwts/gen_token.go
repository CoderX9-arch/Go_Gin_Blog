package jwts

import (
	"blog/global"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

// GenToken创建Token
// 生成token
func GenToken(user JwtPayLoad) (string, error) {
	MySecret = []byte(global.Config.Jwy.Secret)
	claim := CustomClaims{
		user,
		jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Hour * time.Duration(global.Config.Jwy.Expires))),
			Issuer:    global.Config.Jwy.Issuer,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)
	return token.SignedString(MySecret)
}
