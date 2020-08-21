package test

import (
	"fanatic/router"
	"github.com/gin-gonic/gin"
)

func setupApp() *gin.Engine {
	app := gin.Default()
	router.LoadRouter(app)

	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"author": "welong",
		})
	})

	return app
}
