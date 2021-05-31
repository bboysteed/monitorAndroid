package monkey

import (
	"fmt"
	"monitorAndroid/models"
	"monitorAndroid/utils"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

//后期加入，多设备时
type monkeyEngin struct {
}

func RunningMonkey(um *models.UpMission, dm *models.DownMission, logPath string) {
	fmt.Printf("mission start at:%v\n", time.Now().Format("2006/01/02 15:04:05"))
	for i := 1; i <= um.Thread; i++ {
		now := time.Now().Format("2006_01_02_15_04_05")
		//path,_:=filepath.Abs(logPath)
		storePath := filepath.Join(logPath, now)

		err := os.MkdirAll(storePath, os.ModePerm)
		if err != nil {
			fmt.Printf("CREAT DIR path with err :%v\n", err)
		}
		for _, app := range um.Apps {
			go RunAMonkey(app, strconv.Itoa(um.Interval), storePath, dm)
			time.Sleep(time.Second)
		}
	}
}

func RunAMonkey(app, interval, storePath string, dm *models.DownMission) {
	storelogPath := filepath.Join(storePath, app+".log")
	fmt.Printf("开始测试软件:%v...\n", app)
	command := fmt.Sprintf("adb shell monkey -p %v --pct-syskeys 0 --throttle %v -v -v -v 10 >%v", app, interval, storelogPath)
	fmt.Printf("test command is:%v\n", command)

	_, err := utils.Exec(command)
	if err != nil {
		fmt.Printf("monkey command exec failed,err is :%v\n", err)
	}
	//cmd := exec.Command("bash","-c",command)
	//err:= cmd.Start()
	//if err != nil {
	//	fmt.Printf("monkey start failed,err is :%v\n",err)
	//}else {
	//	err = cmd.Wait()
	//	if err != nil {
	//		fmt.Printf("command finished with err:%v\n",err)
	//	}
	//}
	dm.Progress++

}
