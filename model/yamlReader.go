package model

import (
	"io/ioutil"

	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	// Env struct {
	// 	AgentUrl string `yaml: "agent"`
	// }
	Agent struct {
		Url string `yaml: "url"`
	}
	Redis struct {
		Endpoint string   `yaml: "endpoint"`
		Url      string   `yaml: "url"`
		Port     int      `yaml: "port"`
		Password string   `yaml: "password"`
		Section  []string `yaml: "section"`
	}
}

// 讀取 conf/monitor.yaml 並回傳 Config Instance
func GetConfig() *Config {
	config := Config{}
	file, err := ioutil.ReadFile("conf/monitor.yaml")

	if err != nil {
		log.Fatalf("讀取 yaml 檔發生錯誤", err)
	}

	err = yaml.Unmarshal(file, &config)
	if err != nil {
		log.Fatalf("轉換 yaml 檔發生錯誤", err)
	}
	return &config
}
