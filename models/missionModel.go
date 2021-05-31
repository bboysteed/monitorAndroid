package models

type UpMission struct {
	Name     string   `json:"name"`
	Apps     []string `json:"apps"`
	Thread   int      `json:"thread"`
	Interval int      `json:"interval"`
}
type DownMission struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	State         string `json:"state"`
	StartTime     string `json:"startTime"`
	Progress      int    `json:"progress"`
	DeviceOffline bool   `json:"deviceoffline"`
}

//func (um *UpMission) LoadPostValue(form url.Values) {
//
//	um.Name = form.Get("name")
//	um.Apps = form.Get("apps")
//	um.Thread = form.Get("thread")
//	um.Interval = form.Get("interval")
//}

/*

if __name__ == '__main__':
    print("mission start at ", time.strftime("%Y-%m-%d %H:%M:%S", time.localtime()))
    dirPath = os.path.join(os.path.abspath("./"), datetime.datetime.now().strftime('%m_%d_%H_%M_%S'))
    os.makedirs(dirPath)
    for i in range(1):
        storePath = os.path.join(dirPath, str(i))
        os.makedirs(storePath)
        f = open("allPackages.txt", "r")
        packages = f.readlines()
        random.shuffle(packages)
        for name in packages:
            print(f'\033[32m正在测试软件:{name.strip()}...\033[0m')
            logPath = os.path.join(storePath,name.strip()+".log")
            monkey_count = 100 if name.strip() == "com.chery.launcher" else 1000
            subprocess.run(
                f"adb shell monkey -p {name.strip()} --pct-syskeys 0 --throttle 100 -v -v -v {monkey_count} >{logPath}",
                shell=True)
            sleep(3)


    # print(f'\033[32m正在测试软件:{name.strip()}...\033[0m')

    print("mission end at ", time.strftime("%Y-%m-%d %H:%M:%S", time.localtime()))

*/
