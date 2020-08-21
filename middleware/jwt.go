package middleware

import (
	"fanatic/lib/erro"
	"fanatic/lib/redis"
	"fanatic/lib/token"
	"fanatic/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

var redisService = redis.NewRedis()

func JwtHandle(c *gin.Context) {
	BearerToken := c.GetHeader("authorization")
	if BearerToken == "" {
		c.Error(erro.Unauthorized)
		c.Abort()
	}
	tokenStr := strings.Replace(BearerToken, "Bearer ", "", 1)
	claims, e := token.VerifyAccessToken(tokenStr)
	if e != nil {
		c.Error(e)
		c.Abort()
	} else if claims != nil {
		id, ok := claims["identity"].(float64)
		if !ok {
			c.Abort()
			c.JSON(http.StatusBadRequest,
				erro.Unauthorized.SetMsg("token不合法").SetUrl(c.Request.URL.String()))
		}
		userService := service.NewUserService()
		user := userService.FindOneById(int32(id))
		if user.Account == "" {
			c.Error(erro.Unauthorized.SetMsg("非法操作！"))
			c.Abort()
		}
		key := hasRedis(user.ID)
		if !validateRedis(key, c) {
			rdsKey := fmt.Sprintf("user:%d:permission", user.ID)
			if !hasPermission(int32(user.ID), rdsKey, c) {
				c.Error(erro.Forbidden.SetMsg("权限不足"))
				c.Abort()
			}
		}
		c.Set("id", user.ID)
		c.Next()
	}

}

func hasPermission(uid int32, key string, c *gin.Context) bool {
	permission, exist := c.Get("permission")
	// 若permission不存在,直接跳出权限校验
	if !exist {
		return true
	}
	// 若permission存在,校验开始

	p := service.NewPermissionService()
	userInfo := p.FindUserPermission(uint(uid))
	if len(userInfo.Permission) == 0 {
		return false
	}
	for i := 0; i < len(userInfo.Permission); i++ {
		redisService.Lpush(key, userInfo.Permission[i].Name)
	}
	for i := 0; i < len(userInfo.Permission); i++ {
		if userInfo.Permission[i].Name == permission {
			return true
		}
	}
	return false
}

func hasRedis(uid uint) string {
	userRedisKey := fmt.Sprintf("user:%d:permission", uid)
	exists, err := redisService.Exists(userRedisKey)
	if err != nil {
		log.Println(err)
	}
	if exists > 0 {
		return userRedisKey
	}
	return ""
}

func validateRedis(key string, c *gin.Context) bool {
	if key != "" {
		var i int64
		length, _ := redisService.Llen(key)
		permission, _ := c.Get("permission")
		for i = 0; i < length; i++ {
			// redis中存有用户权限，则在此校验
			redisPermission, _ := redisService.Lindex(key, i)
			if redisPermission == permission {
				return true
			}
		}
	}
	return false
}
