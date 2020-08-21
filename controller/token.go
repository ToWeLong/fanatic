package controller

import (
	"fanatic/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

var tokenService = service.NewTokenService()

// 通过refreshToken换取accessToken
func RefreshToAccess(c *gin.Context) {
	type Refresh struct {
		RefreshToken string `json:"refresh_token" binding:"required"`
	}
	var refresh Refresh

	if err := c.ShouldBindJSON(&refresh); err != nil {
		c.Error(err)
		return
	}
	access,err := tokenService.RefreshExchangeAccessToken(refresh.RefreshToken)
	if err!=nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"access_token": access,
	})
}
