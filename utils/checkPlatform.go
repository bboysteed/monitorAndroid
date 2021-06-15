package utils

import (
	"log"
	"os"
	"runtime"
)

func CheckPlatform(){
	platform := runtime.GOOS
	if platform!="windows"{
		log.Printf("this operation system is not supported!\nonly supported windows!")
		os.Exit(1)
	}
}
