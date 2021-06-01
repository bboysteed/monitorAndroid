package models

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"monitorAndroid/utils"
	"regexp"
)

type Phone struct {
	AndroidID string   `json:"androidID"`
	Packages  []string `json:"packages"`
}

func (p *Phone) GetAllPackages(phones *[]Phone) {
	for _, ph := range *phones {
		if p.AndroidID == ph.AndroidID {
			p.Packages = ph.Packages
			return
		}
	}
	var out string
	var err error
	//var mac string
	//get mac adress
	//p.GetMacAddress()

	//get packages
	res := make([]string, 0)
	out, err = utils.Exec("adb shell pm list packages")
	if err != nil {
		log.Printf("list package 命令失败，err is：%v\n", err.Error())
	} else {
		reg := regexp.MustCompile(`package:(.*?)\s+`)
		match := reg.FindAllStringSubmatch(out, -1)
		//fmt.Printf("match is : %v", match)

		for _, v := range match {
			res = append(res, v[1])
		}
	}

	p.Packages = res
	//fmt.Printf("phone now is : %v\n",p)
	//新设备存入json
	*phones = append(*phones, *p)

	newJson, marshallErr := json.Marshal(phones)
	if marshallErr != nil {
		log.Printf("marshall phones err, err is:%v\n", marshallErr)
	} else {
		err = ioutil.WriteFile("./packages.json", newJson, 0644)
		if err != nil {
			log.Printf("Write json err,err is: %v\n", err)
		}
	}

}

func (p *Phone) GetMacAddress() {
	var out string
	var err error
	var androidID string
	//get mac adress
	//adb shell settings get secure android_id
	//adb shell cat /sys/class/net/wlan0/address
	out, err = utils.Exec("adb shell cat /sys/class/net/wlan0/address")
	if err != nil {
		log.Printf("cat mac adress 命令失败，err is：%v\n", err.Error())
	} else {
		reg := regexp.MustCompile(`\s+`)
		androidID = reg.ReplaceAllString(out, "")
		//mac = strings.ReplaceAll(string(out),"\r\n","")
	}
	p.AndroidID = androidID
}
