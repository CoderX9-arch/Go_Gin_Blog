package main

import (
	"blog/core"
	"blog/global"
	"blog/utils/jwts"
	"fmt"
)

func main() {
	core.InitConf()
	fmt.Println(global.Config)

	token, err := jwts.GenToken(jwts.JwtPayLoad{
		UserID:   "1",
		Role:     "1",
		Username: "test",
		Nickname: "xxx",
	})
	fmt.Println(token, err)

	claims, err := jwts.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6InRlc3QiLCJuaWNrbmFtZSI6Inh4eCIsInJvbGUiOiIxIiwidXNlcmlkIjoiMSIsImV4cCI6MTczMzM4Njc4NS45MDA4NjgyfQ.TcMrNxUzhMsKlhpg6v13WzcRWMUlc2he0gP2M76AVXo")
	fmt.Println(claims, err)
}
