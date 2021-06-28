package models

import (
	"github.com/jinzhu/configor"
	"log"
)

type CFG struct {
	OsName         string
	OsArch         string
	Host           string
	GetMacCommand  string
	GetTemperature string
}

func (cfg *CFG) LoadYAML() {
	err := configor.Load(cfg, "./conf.yml")
	if err != nil {
		log.Printf("conf load failed, err is: %v\n", err)
	}
	log.Printf("cfg is: %#v\n", cfg)
}
