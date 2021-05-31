package main

import (
	"fmt"
	"os/exec"
)

func main() {
	//str := `User 3%, System 9%, IOW 0%, IRQ 0%`
	//	str := `Total RAM: 2004576 kB
	// Free RAM: 962859 kB (28123 cached pss + 151424 cached + 783312 free)
	// Used RAM: 832478 kB (791862 used pss + 6700 buffers + 8880 shmem + 25036 slab)
	//`
	//	resList := utils.RegFind(`Total\s+RAM:\s+(\d+)\s+kB`,str,1)
	//	resList := utils.RegFind(`Used\s+RAM:\s+(\d+)\s+kB`,str,1)
	//
	//	//cpuRate :=
	//	//fmt.Println()
	//	fmt.Println(resList)

	//out, err := exec.Command("bash","-c","adb shell monkey -p com.chery.eol -v 1000").Output()
	//if err != nil {
	//	fmt.Printf("top命令失败，err is：%v", err.Error())
	//
	//} else {
	//	fmt.Println(string(out))
	//}
	//fmt.Println(time.Now().Format("06/01/02 15:04"))
	//
	//netInterfaces, err := net.Interfaces()
	//if err != nil {
	//	fmt.Println("net.Interfaces failed, err:", err.Error())
	//
	//}
	//
	//for i := 0; i < len(netInterfaces); i++ {
	//	if (netInterfaces[i].Flags & net.FlagUp) != 0 {
	//		addrs, _ := netInterfaces[i].Addrs()
	//
	//		for _, address := range addrs {
	//			if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
	//				if ipnet.IP.To4() != nil {
	//					fmt.Println(ipnet.IP.String())
	//
	//
	//				}
	//			}
	//		}
	//	}
	//}
	//command:=fmt.Sprintf("adb shell monkey -p com.redstone.fotaapp --pct-syskeys 0 --throttle 400 -v -v -v 10 >/home/steed/go_projects/monitorAndroid/logs/2021_05_29_22_20_26/com.redstone.fotaapp.log")
	//out,_ := exec.Command("CMD","/C", "WHOAMI").Output()
	//fmt.Println(string(out))
	//fmt.Println(runtime.GOARCH)
	//fmt.Println(runtime.GOOS)

	out, err := exec.Command("CMD", "/C", "start adb shell whoami>a.txt").Output()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(out))

}
