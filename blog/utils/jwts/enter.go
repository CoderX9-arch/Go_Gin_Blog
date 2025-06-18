package jwts

import (
	"github.com/dgrijalva/jwt-go/v4"
)

type JwtPayLoad struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Role     int    `json:"role"`
	UserID   string `json:"userid"`
}

var MySecret []byte

type CustomClaims struct {
	JwtPayLoad
	jwt.StandardClaims
}
