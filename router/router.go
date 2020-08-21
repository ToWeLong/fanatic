package router

import (
	"fanatic/controller"
	"fanatic/lib/erro"
	"fanatic/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func LoadRouter(app *gin.Engine, middle ...gin.HandlerFunc) *gin.Engine {

	// handle no method
	app.NoMethod(func(c *gin.Context) {
		s := erro.NoMethodMatched.SetUrl(c.Request.URL.String())
		c.JSON(http.StatusForbidden, s)
	})

	// handle no router
	app.NoRoute(func(c *gin.Context) {
		s := erro.NoRouteMatched.SetUrl(c.Request.URL.String())
		c.JSON(http.StatusNotFound, s)
	})

	app.Use(middleware.ErrorHandle)
	app.Use(middleware.CORS)
	// 方便在load router之后注册中间件
	app.Use(middle...)
	user := app.Group("/user")
	user.POST("/login", controller.Login)
	user.POST("/register", controller.Register)
	user.POST("/permission", controller.RegisterUserPermission)
	user.PUT("/permission", controller.EditUserPermission)
	user.DELETE("/permission", controller.DeleteUserPermission)
	user.GET("/permission/all", controller.PermissionAll)
	user.POST("/token", controller.RefreshToAccess)
	user.GET("/",
		middleware.JwtHandle,
		controller.FindUser,
	)

	book := app.Group("/book")
	book.GET("/register",
		middleware.RouteMeta("添加图书", "图书", 1),
		controller.RegisterBook,
	)
	book.GET("/all", controller.All)
	book.DELETE("/delete",
		middleware.RouteMeta("删除图书", "图书", 1),
		middleware.JwtHandle,
		controller.DeleteBook,
	)
	return app
}
