package models

import (
	"fmt"
	"monitorAndroid/utils"
	"time"
)

type Ram struct {
	Rate int `json:"rate"`
}

func (r *Ram) RefreshRate() {

	out, err := utils.Exec("adb shell dumpsys meminfo")
	if err != nil {
		fmt.Printf("dumpsys meminfo命令失败，err is：%v\n", err.Error())
	} else {
		total := utils.RegFind(`Total\s+RAM:\s+(\d+)\s+kB`, out, 1)
		used := utils.RegFind(`Used\s+RAM:\s+(\d+)\s+kB`, out, 1)
		r.Rate = int((float32(utils.Str2Uint(used[0])) / float32(utils.Str2Uint(total[0]))) * 100.0)

	}
	time.Sleep(time.Second)
}
