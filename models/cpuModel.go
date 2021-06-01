package models

import (
	"log"
	"monitorAndroid/utils"
	"regexp"
	"time"
)

type Cpu struct {
	UserRate    int     `json:"user_rate"`
	SystemRate  int     `json:"system_rate"`
	AllRate     int     `json:"all_rate"`
	Temperature float32 `json:"temperature"`
}

func (c *Cpu) RefreshRate() {
	out, err := utils.Exec("adb shell top -n 1 -d 1")
	if err != nil {
		log.Printf("top命令失败，err is：%v\n", err.Error())
	} else {
		matchs := utils.RegFind(`User\s+(\d+)%.*?System\s+(\d+)%`, out, 1, 2)
		if len(matchs) < 2 {
			panic("this android platform is not suitable")
		}
		c.UserRate = utils.Str2Uint(matchs[0])
		c.SystemRate = utils.Str2Uint(matchs[1])
		c.AllRate = c.UserRate + c.SystemRate
	}
	out, err = utils.Exec("adb shell cat /sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		log.Printf("获取温度命令失败，err is：%v\n", err.Error())
	} else {
		c.Temperature = float32(utils.Str2Uint(regexp.MustCompile(`\s*`).ReplaceAllString(out,""))) / 1000.0
		log.Printf("temperature is : %v", c.Temperature)
	}
	time.Sleep(time.Second)
}
