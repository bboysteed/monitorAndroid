package main

import (
	"encoding/json"
	"github.com/gorilla/websocket"
	"io/ioutil"
	"log"
	"monitorAndroid/models"
	"monitorAndroid/monkey"
	"monitorAndroid/utils"
	"net/http"
	"os"
	"strings"
	"time"
)


var (
	cfg = &models.CFG{}
	cpu = &models.Cpu{}
	ram = &models.Ram{}
	mgs = &models.Message{
		Cpu: cpu,
		Ram: ram,
	}
	phoneNow = models.Phone{}
	phones   []models.Phone

	upmission   = &models.UpMission{}
	downmission = &models.DownMission{}
	logPath     = "./logs"

	upgrader = websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

//adb shell cat /sys/class/net/wlan0/address
func main() {
	cfg.LoadYAML()

	initSome()


	http.Handle("/", http.FileServer(http.Dir("dist/")))
	http.HandleFunc("/ws", HandleWebsocket)
	http.HandleFunc("/getAppList", handleGetApps)
	http.HandleFunc("/getApisAndID", handleGetApis)
	http.HandleFunc("/commitMission", handleCommitMission)
	http.HandleFunc("/missionStatews", handleMissionStatews)

	log.Println("Listening on localhost:8089")
	log.Fatal(http.ListenAndServe(":8089", nil))
}

func handleMissionStatews(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("handleMissionStatews ws upgrade filed,err is: %v", err)
		return
	}

	for downmission.State == "running"{
		log.Println(downmission)
		if downmission.Progress >= upmission.Thread*len(upmission.Apps) {
			downmission.State = "finished"
		} else {
			downmission.State = "running"
		}

		conn.WriteJSON(downmission)
		time.Sleep(3 * time.Second)
		if downmission.DeviceOffline{
			break
		}
	}

	//downmission.State = "running"
	downmission.Progress = 0
}

func handleCommitMission(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Access-Control-Allow-Origin", "*")
	writer.Header().Set("Access-Control-Allow-Credentials", "true")
	writer.Header().Set("Access-Control-Allow-Headers", "Content-Type,Access-Control-Allow-Headers,Content-Length,Accept,Authorization,X-Requested-With")
	writer.Header().Set("Access-Control-Allow-Methods", "PUT,POST,GET,DELETE,OPTIONS")

	if downmission.DeviceOffline {
		writer.Write([]byte(`{"msg":"offline"}`))
		return
	}
	if downmission.State == "running" {
		writer.Write([]byte(`{"msg":"running"}`))
		return
	}
	// 检查是否POST请求
	if request.Method != "POST" {
		writer.WriteHeader(405)
		return
	}
	//body解析
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Printf("read from http body with err:%v\n", err)
	}
	//log.Printf("%s",body)
	json.Unmarshal(body, upmission)
	log.Printf("upmission:%#v\n", upmission)
	downmission.ID = 1
	downmission.Name = upmission.Name
	downmission.StartTime = time.Now().Format("06/01/02 15:04")
	downmission.State = "running"
	downmission.DeviceOffline = false
	rsp, _ := json.Marshal(downmission)
	//fmt.Println(string(rsp))
	writer.Write(rsp)
	go monkey.RunningMonkey(upmission, downmission, logPath)

}

func handleGetApis(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	StaticMap := make(map[string]string)
	StaticMap["wsAddress"] = strings.Join([]string{"ws://", cfg.Host, "/ws"}, "")
	StaticMap["AppsAddress"] = strings.Join([]string{"https://", cfg.Host, "/getApps"}, "")
	StaticMap["androidID"] = phoneNow.AndroidID

	rsp, _ := json.Marshal(StaticMap)
	w.Write(rsp)
}

func initSome() {
	setplatforArch()
	utils.CheckPlatform()
	utils.CheckRegister()
	checkConnection()
	loadJson()
	phoneNow.GetMacAddress()
	cfg.LoadYAML()
	go checkDeviceOffline()
}

func setplatforArch() {
	cpu.Arch = cfg.OsArch
	ram.Arch = cfg.OsArch
}

func checkDeviceOffline() {
	for {
		_, err := utils.Exec("adb shell ls /data") //
		if err != nil {
			log.Printf("devices is offline,please reconnect device...\n")
			downmission.DeviceOffline = true
			downmission.State="finished"
		} else {
			if downmission.DeviceOffline==true{
				log.Println("device reconnected!!")
			}
			downmission.DeviceOffline = false
		}
		time.Sleep(10 * time.Second)
	}

}

func checkConnection() {
	var out string
	var err error
	out, err = utils.Exec("adb shell ls /data") //
	if err != nil {
		log.Printf("devices is offline,please connect device,then try again...,err is:%v", err)
		os.Exit(1)
	} else {
		downmission.DeviceOffline = false
		log.Println("设备在线:")
		out, _ = utils.Exec("adb devices")
		log.Printf(out)
	}
}

func loadJson() {
	bytes, err := ioutil.ReadFile("./packages.json")
	if err != nil {
		log.Println("读取packages.json文件失败", err)
		return
	}

	err = json.Unmarshal(bytes, &phones)
	if err != nil {
		log.Println("解析packages.json文件失败", err)
		return
	}
	//fmt.Printf("%v\n", phones)
}

func HandleWebsocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("ws upgrade filed,err is: %v", err)
		return
	}
	for !downmission.DeviceOffline {
		go cpu.RefreshRate()
		ram.RefreshRate()
		conn.WriteJSON(mgs)

	}
	conn.WriteJSON(&struct {
		Msg string `json:"msg"`
	}{
		Msg: "offline",
	})

}

func handleGetApps(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	if downmission.DeviceOffline {
		w.Write([]byte(`{"msg":"offline"}`))
		return
	}
	log.Printf("req url is: %v\n", req.URL)
	phoneNow.GetAllPackages(&phones)
	rep, _ := json.Marshal(phoneNow)
	w.Write(rep)

}
