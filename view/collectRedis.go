package view

import (
	"RedisMonitor/controller"
	"encoding/json"
	"fmt"
	"log"
)

// 蒐集 Redis 系統資料方法
// 組合 controller 中的方法
func CollectRedis() {
	log.Println("開始蒐集 Redis 資訊")
	log.Println("===========================")
	client := controller.GetRedisClient()
	sessions := controller.GetConfigData().Redis.Section
	for _, session := range sessions {
		log.Println(fmt.Sprintf("抓取 Redis Session %s", session))
		infoStr := controller.GetRedisInfo(client.Info(session).String())
		infoJson, err := json.Marshal(&infoStr)
		if err != nil {
			log.Fatalf(fmt.Sprintf("轉換 Session %s 至 Json 格式時發生錯誤，錯誤訊息為 %s", err))
		}
		controller.PushData(string(infoJson))
	}
	log.Println("===========================")
	log.Println("結束程式")
}
