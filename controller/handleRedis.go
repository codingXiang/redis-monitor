package controller

import (
	"log"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// 取得 Redis Info
func GetRedisInfo(infoStr string) []*Data {
	datas := strings.Split(infoStr, "\n")
	result := []*Data{}
	now := time.Now().Unix()
	for _, content := range datas {
		data := parseRedisInfo(content, now)
		if data != nil {
			result = append(result, data)
		}
	}
	return result
}

// 解析 Redis Info 字串並回傳 Data
func parseRedisInfo(content string, now int64) *Data {
	tmp := strings.Split(content, ":")
	if len(tmp) > 1 {
		return GetPushAgentData(ConfigData.Redis.Endpoint, tmp[0], stringToFloat(tmp[1]), now)
	}
	return nil
}

// 去除非數值的參數
func removeNoiseStr(str string) string {
	re := regexp.MustCompile(`[-]?\d[\d,]*[\.]?[\d{2}]*`)
	result := re.FindAllString(str, 1)
	if len(result) > 0 {
		return result[0]
	} else {
		return ""
	}
}

// 將 String 轉換成 Float
func stringToFloat(str string) float64 {
	_str := removeNoiseStr(str)
	if _str == "" {
		return 0
	} else {
		value, err := strconv.ParseFloat(_str, 64)
		if err != nil {
			log.Fatal("字串轉換 Int 發生錯誤，%s", err)
		}
		return value
	}

}
