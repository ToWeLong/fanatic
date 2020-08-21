package test

import (
	"fanatic/lib/token"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"testing"
	"time"
)

func TestJwt(t *testing.T)  {

	viper.Set("token.secretKey","CCbsch4j3b1XLaME")
	expireTime := time.Now().Add(time.Second * 120)
	tokenStr := token.GenerateToken(token.ACCESS,2,expireTime)
	log.Println(tokenStr)
	//time.Sleep(time.Second * 5)
	claims, e := token.VerifyAccessToken(tokenStr)
	if e!=nil{
		log.Info(e)
	}
	log.Println(claims)
	//BearerToken := "Bearer eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpZGVudGl0eSI6MSwic2NvcGUiOiJsaW4iLCJ0eXBlIjoiYWNjZXNzIiwiZXhwIjoxNTk3MTIzMzE5fQ.wbi2kJusq-WLa8csxYczJOS_oLRJVRuo6xuengmCXUQ"
	//tk := strings.Replace(BearerToken,"Bearer ","",1)
	//fmt.Println(tk)
}
