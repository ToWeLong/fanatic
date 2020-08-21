package crypt

import (
	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// @password: 原密码
// return 加密之后的密码
func EncodePassword(psw string)[]byte {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(psw), 5)
	if err != nil {
		log.Info("密码加密失败:",err)
	}
	return hashedPassword
}


// return: 是否正确
func VerfifyPsw(rawPsw, enCodePassword []byte) bool {
	return bcrypt.CompareHashAndPassword(enCodePassword, rawPsw) == nil
}