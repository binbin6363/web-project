package config

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Email  *Email  `yaml:"email"`  // 邮箱配置
	Server *Server `yaml:"server"` // 邮箱配置
	Db     *Db     `yaml:"mysql"`  // 邮箱配置
	Alert  *Alert  `yaml:"alert"`  // 告警配置
}

type Alert struct {
	Addr   string `yaml:"addr"`   // 地址
	Secret string `yaml:"secret"` // Secret
}
type Server struct {
	Addr string `yaml:"addr"` // 服务监听地址
}
type Db struct {
	Addr string `yaml:"addr"` // 服务监听地址
}

type Email struct {
	From     string `yaml:"from"`  // 发送者邮箱
	NickName string `yaml:"nick"`  // 发送者昵称
	Token    string `yaml:"token"` // 邮箱授权码。AXNJPCHCMCMVVKDW
	SmtpHost string `yaml:"host"`  // 邮箱服务器地址
	SmtpPort int    `yaml:"port"`  // 邮箱服务器端口
}

var c Config

func GetConfig() *Config {
	return &c
}

func Init(cfgFile string) {
	yamlFile, err := ioutil.ReadFile(cfgFile)
	if err != nil {
		log.Fatal(err.Error())
	}

	err = yaml.Unmarshal(yamlFile, &c)

	if err != nil {
		log.Fatal(err.Error())
	}
	log.Printf("load config ok, %s", cfgFile)
}
