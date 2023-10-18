package main

import (
	"github.com/spf13/viper"
	"myGoIm/models"
	"myGoIm/router"
	"myGoIm/utils"
	"time"
)

func main() {
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	InitTimer()
	r := router.Router() // router.Router()
	r.Run(viper.GetString("port.server"))
}

// InitTimer 初始化定时器
func InitTimer() {
	utils.Timer(time.Duration(viper.GetInt("timeout.DelayHeartbeat"))*time.Second, time.Duration(viper.GetInt("timeout.HeartbeatHz"))*time.Second, models.CleanConnection, "")
}
