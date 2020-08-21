package main

import (
	"errors"
	"fanatic/config"
	"fanatic/lib/redis"
	validate "fanatic/lib/validator"
	"fanatic/model"
	"fanatic/router"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func main() {

	godotenv.Load()

	if err := config.Init(); err != nil {
		log.Error(err)
	}

	log.SetFormatter(&log.TextFormatter{
		ForceColors:   true,
		FullTimestamp: true,
	})

	model.Init()
	model.Sync()
	model.OpenLog()
	defer model.Close()

	redis.Init()
	defer redis.Close()

	app := gin.Default()
	app.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"api-version": "v1",
			"author":      "welong",
		})
	})

	// set console.log color
	gin.ForceConsoleColor()
	binding.Validator = validate.New()

	router.LoadRouter(app)
	if e := app.Run(viper.GetString("port")); e != nil {
		errors.New("服务器未知错误")
		log.Error(e)
	}

}
