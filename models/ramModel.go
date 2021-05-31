package models

import (
	"fmt"
	"monitorAndroid/utils"
	"os/exec"
	"time"
)

type Ram struct {
	Rate int `json:"rate"`
}

func (r *Ram) RefreshRate() {

	out, err := exec.Command("bash", "-c", "adb shell dumpsys meminfo").Output()
	if err != nil {
		fmt.Printf("dumpsys meminfo命令失败，err is：%v\n", err.Error())
	} else {
		total := utils.RegFind(`Total\s+RAM:\s+(\d+)\s+kB`, string(out), 1)
		used := utils.RegFind(`Used\s+RAM:\s+(\d+)\s+kB`, string(out), 1)
		r.Rate = int((float32(utils.Str2Uint(used[0])) / float32(utils.Str2Uint(total[0]))) * 100.0)

	}
	time.Sleep(time.Second)
}
