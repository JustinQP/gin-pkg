package utils

import (
	"golang.org/x/crypto/bcrypt"
)

func PwdEncode(pwd string) string {
	h, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		return ""
	}

	return string(h)
}

func PwdVerify(pwd, loginPwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(pwd), []byte(loginPwd))
	if err != nil {
		return false
	} else {
		return true
	}
}
