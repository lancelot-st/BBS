package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var Instance *Config

type Config struct {
	Env        string `yaml:"Env"`        //基础环境
	BaseUrl    string `yaml:"BaseUrl"`    //url
	ShowSql    bool   `yaml:"ShowSql"`    //是否打印Sql
	StaticPath string `yaml:"StaticPath"` //静态文件路径

	app struct {
		name string `yaml:"name"`
		Port string `yaml:"Port"`
	} `yaml:"app"`

	log struct {
		level       string `yaml:"level"`
		LogFile     string `yaml:"LogFile"`  //日志存放的文件名
		max_size    int64  `yaml:"max_size"` //最大储存容量
		max_age     int64  `yaml:"max_age"`  //最多储存天数
		max_backups int64  `yaml:"max_backups"`
	} `yaml:"log"`
	//数据库配置
	DB struct {
		Url          string `yaml:"Url"`
		MaxIdleConns int    `yaml:"MaxIdleConns"`
		MaxOpenConns int    `yaml:"MaxOpenConns"`
	} `yaml:"DB"`

	Redis struct {
		host      string `yaml:"host"`
		port      int    `yaml:"port"`
		password  string `yaml:"password"`
		db        int    `yaml:"db"`
		pool_size int    `yaml:"pool_Size"`
	}
}

func Init(fileName string) *Config {
	Instance = &Config{}
	yamlFile, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
	errYaml := yaml.Unmarshal(yamlFile, Instance)
	if errYaml != nil {
		log.Fatal(errYaml)
	}
	return Instance
}
