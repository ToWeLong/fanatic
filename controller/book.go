package controller

import (
	"fanatic/service"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

type Book struct {
	Title string `form:"title" binding:"min=4,max=20"`
	Author string `form:"author" binding:"required"`
}

func RegisterBook(c *gin.Context) {
	var b Book
	if err := c.ShouldBindWith(&b, binding.Query); err != nil {
		_ = c.Error(err)
		return
	}
	value,_ := c.Get("permission")
	c.JSON(http.StatusOK,gin.H{
		"permission":value,
	})
}

func DeleteBook(c *gin.Context)  {
	p := service.NewPermissionService()
	userInfo := p.FindUserPermission(1)
	c.JSON(http.StatusOK,userInfo)
}

func All(c *gin.Context)  {
	p := service.NewPermissionService()
	c.JSON(200,p.FindAllPermission())
}
