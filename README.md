# Redis 監控工具
此 Redis 監控工具目前僅支援 Open-Falcon，因此環境中必須要有 Open-Falcon 的 Agent 資料才能順利儲存

## 環境需求
- 系統：MacOS / Linux
- 開發語言：Golang
- 相依 Package
  - gopkg.in/yaml.v2
  - gopkg.in/redis.v4

## 指標
目前採集的指標是使用 Info Section 的方式，Section 的指定設定在 conf/monitor.yaml 當中

## 部署
### 實機部署
1. 使用 make file
````
make build
make pack
````
2. 將打包好的 ```redis-monitor-v1.0.0.tar.gz``` 移動至 /monitor
3. 執行解壓縮動作
````
cd /monitor
tar xvf redis-monitor-v1.0.0.tar.gz
````
4. 修改 ```/monitor/conf/monitor.yaml``` 資料
5. 透過 crontab 的方式固定時間運行腳本
````
#crontab -e
*/1 * * * * cd /monitor && ./redis-monitor >> /monitor/log
````

### Docker部署
1. 在根目錄使用 docker build 指定進行打包
````
docker build -t redis-monitor:1.0.0 .
````

2. 使用 docker run 指定進行部署
````
docker run --name redis-monitor -d \
-e Agent_Url="127.0.0.1:1988" \
-e Redis_Endpoint="Redis Server" \
-e Redis_Port=6379 \
-e Redis_Password="" \
redis-monitor:1.0.0
````