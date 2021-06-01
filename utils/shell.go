package utils

import (
	"errors"
	"log"
	"os/exec"
	"runtime"
)

func Exec(command string) (string, error) {
	platform := runtime.GOOS

	if platform == "windows" {
		out, err := exec.Command("CMD", "/C", command).Output()
		if err != nil {
			log.Printf("[windows] run command:%v failed with err:%v\n", command, err)
			return "", err
		} else {
			return string(out), nil
		}

	} else if platform == "linux" {
		out, err := exec.Command("bash", "-c", command).Output()
		if err != nil {
			log.Printf("[linux] run command:%v failed with err:%v\n", command, err)
			return "", err
		} else {
			return string(out), nil
		}
	} else {
		return "", errors.New("unsupported platform！")
	}
}
