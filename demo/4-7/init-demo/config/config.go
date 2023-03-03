package config

import (
	"sync"
)

var once sync.Once

type GConfig struct {
	C1 string
	C2 string
	C3 string
}

var GlobalConfig = &GConfig{}

func GetConfig() *GConfig {
	once.Do(func() {
		iniSysConfig()
		iniAppConfig()
	})
	return GlobalConfig
}

func iniSysConfig() {
	GlobalConfig.C1 = "c1"

}

func iniAppConfig() {
	GlobalConfig.C2 = "c2"
	initKey()
}

func initKey() {
	tmpConfig := GetConfig()
	GlobalConfig.C3 = tmpConfig.C1 + "c3"
}
