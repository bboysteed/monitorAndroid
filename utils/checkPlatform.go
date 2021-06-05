package utils

import (
	"log"
	"runtime"
)

func CheckPlatform(){
	platform := runtime.GOOS
	if platform!="windows"{
		log.Printf("this operation system is not supported!\nonly supported windows!")
	}
}
