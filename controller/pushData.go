package controller

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
)

// 上傳至 Agent 的格式
type Data struct {
	Metric      string  `json: "Metric"`
	Endpoint    string  `json: "endpoint"`
	Timestamp   int64   `json: "timestamp"`
	Step        int     `json: "step"`
	Value       float64 `json: "value"`
	CounterType string  `json: "counterType"`
}

// 上傳資料到 Open-Falcon Agent
func PushData(jsonData string) {
	var agentUrl string = GetConfigData().Agent.Url
	log.Println("上傳資料至 ", agentUrl)
	req, err := http.NewRequest("POST", agentUrl, bytes.NewBufferString(jsonData))
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("發送資料至 Agent 時發生錯誤，錯誤為 %s", err)
	}
	defer resp.Body.Close()
	log.Println("上傳完畢，Http Status Code : ", resp.Status)
	fmt.Println()
}
