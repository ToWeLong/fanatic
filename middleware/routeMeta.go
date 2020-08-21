package middleware

import (
	"fanatic/service"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var permissionService = service.NewPermissionService()

func RouteMeta(permission, module string, mount int) gin.HandlerFunc {
	return func(c *gin.Context) {
		isExist := permissionService.FindPermission(permission)
		if isExist{
			// 若权限不存在，则写入数据库
			if permissionService.RegisterPermission(permission, module, mount){
				c.Set("permission",permission)
				c.Next()
			}else {
				log.Info("权限挂载失败")
				c.Abort()
			}
		}
		// 若权限存在，将permission挂载到上下文
		c.Set("permission",permission)
		c.Next()
	}
}
