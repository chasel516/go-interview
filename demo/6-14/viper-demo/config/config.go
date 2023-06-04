package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"log"
	"time"
)

var Cfg = &Config{}

type Config struct {
	App   App   `yaml:"app"`
	Mysql Mysql `mapstructure:"mysql"`
}

type App struct {
	Env string `yaml:"env"` //yaml标签可以解析
	//Env string `mapstructure: "env"` //不带下划线的配置项，不受:后面空格的影响
	HttpPort int `mapstructure:"http_port"`
	//HttpPort int `mapstructure: "http_port"` //:后有空格不能解析
	//HttpPort int `yaml:"http_port"` //yaml标签不能解析

}

type Mysql struct {
	DBName            string        `mapstructure:"dbname"`
	User              string        `mapstructure:"user"`
	Password          string        `mapstructure:"password"`
	Host              string        `mapstructure:"host"`
	MaxOpenConn       int           `mapstructure:"max_open_conn"`
	MaxIdleConn       int           `mapstructure:"max_idle_conn"`
	ConnMaxLifeSecond time.Duration `mapstructure:"conn_max_life_second"`
	TablePrefix       string        `mapstructure:"table_prefix"`
}

// 加载配置，失败直接panic
func LoadConfig() {
	viper := viper.New()
	//1.设置配置文件路径
	viper.SetConfigFile("config/config.yml")
	viper.SetConfigType("yaml") // 设置配置文件类型

	//2.配置读取
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	//3.将配置映射成结构体
	if err := viper.Unmarshal(&Cfg); err != nil {
		logrus.Error(err)
		panic(err)
	}

	//4. 监听配置文件变动,重新解析配置
	viper.WatchConfig()

	// 配置文件发生变化时的回调函数
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		if err := viper.Unmarshal(&Cfg); err != nil {
			logrus.Error(err)
		}
		log.Printf("%+v", Cfg)
	})
}
