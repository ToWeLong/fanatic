package controller

import (
	"fanatic/form"
	"fanatic/lib/erro"
	"fanatic/model"
	"fanatic/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	userService = service.NewUserService()
	permissionService = service.NewPermissionService()
)

func Login(c *gin.Context) {
	var (
		login form.UserForm
		err   error
	)
	if err = c.ShouldBindJSON(&login); err != nil {
		c.Error(err)
		return
	}
	succ, access, refresh := userService.UserLogin(login.Account, login.Password)
	if !succ {
		c.Error(erro.UserNotFound)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token":  access,
		"refresh_token": refresh,
	})
}

func Register(c *gin.Context) {
	var (
		register form.UserForm
		err      error
	)
	if err = c.ShouldBindJSON(&register); err != nil {
		c.Error(err)
		return
	}

	isExist := userService.RegisterUser(register.Account, register.Password)
	if isExist {
		c.Error(erro.UserExist)
		return
	}
	c.JSON(http.StatusOK, erro.OK.SetUrl(c.Request.URL.String()))
}

func FindUser(c *gin.Context) {
	var user model.User
	account, err := c.GetQuery("account")
	if !err {
		c.Error(erro.ParamsErr)
		return
	}
	user = userService.FindOneByAccount(account)
	if user.Account == "" {
		c.Error(erro.UserNotFound)
		return
	}
	c.JSON(http.StatusOK, user)
}

// 权限相关的接口
func RegisterUserPermission(c *gin.Context)  {
	var (
		permissionForm form.PermissionForm
		err error
	)
	if err = c.ShouldBindJSON(&permissionForm); err != nil {
		c.Error(err)
		return
	}
	if err = permissionService.RegisterUserPermission(permissionForm.UserId,permissionForm.PermissionId); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, erro.OK.SetUrl(c.Request.URL.String()))
}

func EditUserPermission(c *gin.Context)  {
	var (
		permissionEditForm form.PermissionEditForm
		err error
	)
	if err = c.ShouldBindJSON(&permissionEditForm); err != nil {
		c.Error(err)
		return
	}
	if err =permissionService.EditUserPermission(permissionEditForm.UserId,permissionEditForm.OldPermissionId,permissionEditForm.NewPermissionId); err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, erro.OK.SetUrl(c.Request.URL.String()))
}

func PermissionAll(c *gin.Context)  {
	p := service.NewPermissionService()
	c.JSON(http.StatusOK,p.FindAllPermission())
}

func DeleteUserPermission(c *gin.Context)  {
	var (
		permissionForm form.PermissionForm
		err error
	)
	if err = c.ShouldBindJSON(&permissionForm); err != nil {
		c.Error(err)
		return
	}
	p := service.NewPermissionService()
	if err = p.DeleteUserPermission(permissionForm.UserId,permissionForm.PermissionId);err!=nil{
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, erro.OK.SetUrl(c.Request.URL.String()))
}
