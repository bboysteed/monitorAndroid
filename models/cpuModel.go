package models

import (
	"log"
	"monitorAndroid/utils"
	"regexp"
	"strconv"
	"time"
)

type Cpu struct {
	PlatForm
	UserRate    int     `json:"user_rate"`
	SystemRate  int     `json:"system_rate"`
	AllRate     int     `json:"all_rate"`
	Temperature float32 `json:"temperature"`
}

func (c *Cpu) RefreshRate() {
	if c.Arch=="arm"{
		out, err := utils.Exec("adb shell top -n 1 -d 1")
		if err != nil {
			log.Printf("[arm] top命令失败，err is：%v\n", err.Error())
		} else {
			matchs := utils.RegFind(`User\s+(\d+)%.*?System\s+(\d+)%`, out, 1, 2)
			if len(matchs) < 2 {
				log.Println("[arm] 获取cpu信息失败")
				return
			}
			c.UserRate = utils.Str2Uint(matchs[0])
			c.SystemRate = utils.Str2Uint(matchs[1])
			c.AllRate = c.UserRate + c.SystemRate
		}

	}else if c.Arch=="x86"{
		out, err := utils.Exec("adb shell dumpsys cpuinfo")
		if err != nil {
			log.Printf("[x86] top命令失败，err is：%v\n", err.Error())
		} else {
			matchs := regexp.MustCompile(`([\d.]+)%\s+TOTAL:`).FindStringSubmatch(out)
			if len(matchs) < 2 {
				log.Println("[x86] 获取cpu信息失败")
				return
			}
			f,_:=strconv.ParseFloat(matchs[1],32)
			c.AllRate = int(f)
		}
	}


	//温度不分arch
	out, err := utils.Exec("adb shell cat /sys/class/thermal/thermal_zone0/temp")
	if err != nil {
		log.Printf("获取温度命令失败，err is：%v\n", err.Error())
	} else {
		c.Temperature = float32(utils.Str2Uint(regexp.MustCompile(`\s*`).ReplaceAllString(out,""))) / 1000.0
		log.Printf("temperature is : %v", c.Temperature)
	}
	time.Sleep(time.Second)
}
