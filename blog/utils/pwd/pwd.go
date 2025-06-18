package pwd

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"log"
)

func HsahPwd(pwd string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	return string(hash)
}

func CheckPwd(hashPwd string, pwd string) bool {
	byteHsah := []byte(hashPwd)
	err := bcrypt.CompareHashAndPassword(byteHsah, []byte(pwd))
	fmt.Println(err)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
