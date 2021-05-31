package models

import (
	"fmt"
	"github.com/jinzhu/configor"
)

type CFG struct {
	OsName string
	Host   string
}

func (cfg *CFG) LoadYAML() {
	err := configor.Load(cfg, "./conf.yml")
	if err != nil {
		fmt.Printf("conf load failed, err is: %v\n", err)
	}
	fmt.Printf("cfg is: %v\n", cfg)
}
