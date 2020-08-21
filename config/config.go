package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
}

func Init() error {
	cfg := Config{}
	if err := cfg.initConfig();err!=nil{
		return err
	}
	return nil
}

func (c Config) initConfig() error{
	os.Getenv("home")
	getPwd, _ := os.Getwd()
	viper.AddConfigPath(fmt.Sprintf("%s/conf", getPwd))
	viper.SetConfigName("settings")
	viper.SetConfigType("yml")
	viper.SetEnvPrefix("fanatic")
	viper.AutomaticEnv()
	c.watchConfig()
	err := viper.ReadInConfig()
	if err != nil {
		return err
	}
	return nil
}

func (c Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("Config file changed: %s \n", e.Name)
	})
}
