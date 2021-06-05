package utils

import (
	"golang.org/x/sys/windows/registry"
	"log"
	"os"
	"time"
)

func CheckRegister(){
	duration := uint64(90*24*60*60)  //90days
	now:=uint64(time.Now().Unix())
	// 创建：指定路径的项
	// 路径：HKEY_CURRENT_USER\Software\Hello Go
	key, exists, _ := registry.CreateKey(registry.CURRENT_USER, `SOFTWARE\GoAuthor`, registry.ALL_ACCESS)
	defer key.Close()

	// 判断是否已经存在了
	if exists {
		//log.Println(`键已存在`)
		// 读取
		begin, _, _ := key.GetIntegerValue(`registerAt`)
		//log.Println(now-begin)
		if now-begin>duration{
			log.Println("使用期失效,请联系发行者")
			os.Exit(1)
		}
	} else {
		//log.Println(`新建注册表键`)
		// 写入：64位整形值
		err := key.SetQWordValue(`registerAt`, uint64(time.Now().Unix()))
		if err != nil {
			log.Printf("注册失败，请咨询")
		}
	}

}
