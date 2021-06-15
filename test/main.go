//package main

//func main() {
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

//out, err := exec.Command("CMD", "/C", "start adb shell whoami>a.txt").Output()
//if err != nil {
//	fmt.Println(err)
//}
//fmt.Println(string(out))

package main

import (
	"fmt"
	"time"
)

func main() {
	a := fmt.Sprintf("%v", time.Now().Format("2006-01-02,15:04:05"))
	fmt.Println(a)
	//duration := uint64(30*24*60*60)  //30days
	//now:=uint64(time.Now().Unix())
	//// 创建：指定路径的项
	//// 路径：HKEY_CURRENT_USER\Software\Hello Go
	//key, exists, _ := registry.CreateKey(registry.CURRENT_USER, `SOFTWARE\GoAuthor`, registry.ALL_ACCESS)
	//defer key.Close()
	//
	//// 判断是否已经存在了
	//if exists {
	//	fmt.Println(`键已存在`)
	//	// 读取：字符串
	//	begin, _, _ := key.GetIntegerValue(`registerAt`)
	//	fmt.Println(begin)
	//	if now-begin>duration{
	//		log.Println("使用期失效")
	//		os.Exit(1)
	//	}
	//} else {
	//	fmt.Println(`新建注册表键`)
	//	// 写入：64位整形值
	//	err := key.SetQWordValue(`registerAt`, uint64(time.Now().Unix()))
	//	if err != nil {
	//		log.Printf("注册失败，请咨询")
	//	}
	//}

	// 写入：32位整形值
	//key.SetDWordValue(`32位整形值`, uint32(123456))

	// 写入：字符串
	//key.SetStringValue(`字符串`, `hello`)
	// 写入：字符串数组
	//key.SetStringsValue(`字符串数组`, []string{`hello`, `world`})
	// 写入：二进制
	//key.SetBinaryValue(`二进制`, []byte{0x11, 0x22})



	// 读取：一个项下的所有子项
	//keys, _ := key.ReadSubKeyNames(0)
	//for _, key_subkey := range keys {
	//	// 输出所有子项的名字
	//	fmt.Println(key_subkey)
	//}

	//fmt.Println(time.Now().Unix())

	// 创建：子项
	//subkey, _, _ := registry.CreateKey(key, `子项`, registry.ALL_ACCESS)
	//defer subkey.Close()
	//
	//// 删除：子项
	//// 该键有子项，所以会删除失败
	//// 没有子项，删除成功
	//registry.DeleteKey(key, `子项`)
}

//}
