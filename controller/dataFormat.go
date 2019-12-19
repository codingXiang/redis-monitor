package controller

import model "RedisMonitor/model"

// 初始化方法，取得相關設定檔
func init() {
	ConfigData = model.GetConfig()
}

// 取得 ConfigData 方法
func GetConfigData() *model.Config {
	return ConfigData
}

// 將指標與數值轉換為上傳 Agent 格式的資料
func GetPushAgentData(endpoint string, metric string, value float64, timestamp int64) *Data {
	return &Data{
		Endpoint:    endpoint,
		Metric:      metric,
		Value:       value,
		Step:        60,
		Timestamp:   timestamp,
		CounterType: "GAUGE",
	}
}
